package main

import (
	"errors"
	"iter"
	"log"
	"slices"
	"time"
)

type TaskType uint

const (
	//一小时一个的宝箱
	TPSurpriseBenefit TaskType = iota + 1
	//每天的8个任务
	TPDailyBenefit
	//看3个得10点的任务
	TPVideoRewardTabTaskList
	//更多任务 游戏+30点 等等  暂不支持
	TPMoreRewardTab
)

func Sleep() {
	time.Sleep(25 * time.Second)
}
func DoTask(api *QiDianApi, ttps ...TaskType) error {
	sleep := false
	advMainPage, err := api.AdvMainPage()
	if err != nil {
		return err
	}
	taskList := getAllTask(advMainPage, ttps...)
	log.Printf("一共%d个任务\n", len(taskList))
	taskList = NotFinished(taskList)
	log.Printf("还有%d个任务未完成\n", len(taskList))
	for i, task := range taskList {
		finish, err := doTask(api, &task, &sleep)
		if i != len(taskList)-1 {
			Sleep()
		}
		if err != nil {
			log.Printf("第%d个任务[%s]失败:%v\n", i, task.Desc, err)
			return err
		}
		log.Printf("第%d个任务[%s]成功:%v\n", i, task.Desc, finish)
	}
	return nil
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
			return resp, nil
		}
	//104=前往游戏中心玩游戏10分钟奖励10点币
	//103=前往游戏中心任意一款游戏充值1次奖励30点币
	//121=签到互动多重福利(微博)/登陆携程领积分当钱花
	//222=打开推送通知，次日（24h）后可领取奖励
	case 104, 103, 121, 222:
		return singleTask(api, task.TaskId, sleep)
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
func getAllTask(adv *AdvMainPage, ttps ...TaskType) TaskList {
	taskListAll := TaskList{}
	for _, ttp := range ttps {
		switch ttp {
		case TPVideoRewardTabTaskList:
			taskListAll = append(taskListAll, adv.GetVideoRewardTabTaskList(true)...)
		case TPMoreRewardTab:
			//todo 直接提交没效果 暂不处理
			//taskListAll = append(taskListAll, adv.GetMoreRewardTabTaskList()...)
		case TPDailyBenefit:
			taskListAll = append(taskListAll, adv.GetDailyBenefitTaskList()...)
		case TPSurpriseBenefit:
			surprise := adv.GetSurpriseBenefit()
			isReceived := 0
			if surprise.IntervalTime != 0 {
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
		return task.IsFinished == 0 && task.IsReceived == 0
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
