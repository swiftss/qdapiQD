package qdapi

import (
	"errors"
	"iter"
	"log"
	"slices"
	"sync"
	"time"
)

type TaskType uint

const heartBeatTime = 30

func Sleep() {
	time.Sleep(15 * time.Second)
}

func doPlayGame(api *QiDianApi, nickName string, heartBeatCount int) error {
	for i := 0; i < heartBeatCount; i++ {
		beat, err := api.UrlHeartBeat()
		if err != nil {
			return err
		}
		log.Printf("%s:正在玩游戏,玩了:%d秒.返回值:%v\n", nickName, heartBeatTime*(i+1), beat)
		time.Sleep(time.Second * heartBeatTime)
	}
	return nil
}
func needPlayGame(taskList TaskList, nickName string) int {
	for _, task := range taskList {
		if task.TaskType == TPMoreRewardTabPlayGame {
			log.Printf("%s: 游戏共需要玩%ds,已经玩了%ds\n", nickName, task.Total*heartBeatTime, heartBeatTime*task.Process)
			return task.Total - task.Process
		}
	}
	return 0
}
func DoMoreRewardTab(api *QiDianApi, advMainPage *AdvMainPage) error {
	taskList := getAllTask(advMainPage, TPMoreRewardTab)
	tipName := api.TipName()
	log.Printf("%s:一共%d个更多任务\n", tipName, len(taskList))
	taskList = NotFinished(taskList)
	log.Printf("%s:还有%d个更多任务未完成\n", tipName, len(taskList))
	heartBeatCount := needPlayGame(taskList, tipName)
	if heartBeatCount > 0 {
		err := doPlayGame(api, tipName, heartBeatCount)
		if err != nil {
			return err
		}
	}

	for i, task := range taskList {
		finish, err := doTask(api, &task, nil)
		if err != nil {
			log.Printf("%s:第%d个更多任务[%s]失败:%v\n", advMainPage.Data.NickName, i, task.Desc, err)
			return err
		}
		log.Printf("%s:第%d个更多任务[%s]成功:%v\n", advMainPage.Data.NickName, i, task.Desc, finish)
	}
	return nil
}

func DoWatchVideo(api *QiDianApi, advMainPage *AdvMainPage, ttps ...TaskType) error {
	sleep := false
	tipName := api.TipName()
	taskList := getAllTask(advMainPage, ttps...)
	log.Printf("%s:一共%d个看视频任务\n", tipName, len(taskList))
	taskList = NotFinished(taskList)
	log.Printf("%s:还有%d个看视频任务未完成\n", tipName, len(taskList))
	for i, task := range taskList {
		finish, err := doTask(api, &task, &sleep)
		if i != len(taskList)-1 {
			Sleep()
		}
		if err != nil {
			log.Printf("%s:第%d个任务[%s]失败:%v\n", tipName, i, task.Desc, err)
			return err
		}
		log.Printf("%s:第%d个任务[%s]成功:%v\n", tipName, i, task.Desc, finish)
	}
	return nil
}
func DoTask(api *QiDianApi, ttps ...TaskType) error {
	advMainPage, err := api.AdvMainPage()
	if err != nil {
		return err
	}
	wg := &sync.WaitGroup{}
	index := slices.Index(ttps, TPMoreRewardTab)
	if index >= 0 {
		ttps = slices.Delete(ttps, index, index+1)
		go func(api *QiDianApi, advMainPage *AdvMainPage) {
			wg.Add(1)
			DoMoreRewardTab(api, advMainPage)
			wg.Done()
		}(api, advMainPage)
	}
	err = DoWatchVideo(api, advMainPage, ttps...)
	wg.Wait()
	return err
}
func doTask(api *QiDianApi, task *Task, sleep *bool) (*FinishWatch, error) {
	switch task.TaskType {
	//0=8个日常任务,
	case 0:
		return singleTask(api, task.TaskId, sleep)
	//额外看3次得10点那个
	case 3:
		var resp *FinishWatch
		var err error
		for i := task.Process; i < task.Total; i++ {
			resp, err = singleTask(api, task.TaskId, sleep)
			if err != nil {
				return nil, err
			}
		}
		return resp, nil
	//104=前往游戏中心玩游戏10分钟奖励10点币
	//103=前往游戏中心任意一款游戏充值1次奖励30点币
	//121=签到互动多重福利(微博)/登陆携程领积分当钱花
	//222=打开推送通知，次日（24h）后可领取奖励
	case TPMoreRewardTabPlayGame, 103, 121, 222:
		return singleReceiveTaskReward(api, task.TaskId)
	default:
		return singleTask(api, task.TaskId, sleep)
	}
	return nil, errors.New("未知任务类型")
}

func singleTask(api *QiDianApi, taskID string, sleep *bool) (*FinishWatch, error) {
	if *sleep {
		Sleep()
	} else {
		*sleep = true
	}
	resp, err := api.FinishWatch(taskID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func singleReceiveTaskReward(api *QiDianApi, taskID string) (*FinishWatch, error) {
	resp, err := api.ReceiveTaskReward(taskID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func getAllTask(adv *AdvMainPage, ttps ...TaskType) TaskList {
	taskListAll := TaskList{}
	for _, ttp := range ttps {
		switch ttp {
		case TPVideoRewardTabTaskList:
			taskListAll = append(taskListAll, adv.GetVideoRewardTabTaskList(true)...)
		case TPMoreRewardTab:
			taskListAll = append(taskListAll, adv.GetMoreRewardTabTaskList()...)
		case TPDailyBenefit:
			taskListAll = append(taskListAll, adv.GetDailyBenefitTaskList()...)
		case TPSurpriseBenefit:
			surprise := adv.GetSurpriseBenefit()
			isReceived := 0
			if surprise.IntervalTime[:1] != "-" {
				isReceived = 1
			}
			taskListAll = append(taskListAll, Task{
				TaskId:     surprise.TaskId,
				TaskRawId:  surprise.TaskRawId,
				TaskType:   0,
				Desc:       surprise.Desc,
				IsFinished: surprise.IsFinished,
				IsReceived: isReceived,
			})
		}
	}
	return taskListAll
}
func NotFinished(taskList TaskList) TaskList {
	finished := func(task Task) bool {
		support := !slices.Contains(NotSupportTaskType, task.TaskType)
		return task.IsFinished == 0 && task.IsReceived == 0 && support
	}
	return slices.Collect(Filter(finished, slices.Values(taskList)))
}
func Filter[V any](f func(V) bool, s iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range s {
			if f(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}
