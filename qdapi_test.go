package qdapi

import (
	"fmt"
	"github.com/pzx521521/qdapi/sign"
	"sync"
	"testing"
	"time"
)

var api *QiDianApi

var c = QiDianApiConfig{
	QdInfo:  "WnYe23HKHHbFexRmDMhQphCTehSF/SNEL1SCOpnokmfIUobvKBUfPf9v9OO9vQkVpLUc2ZGEfF+RYPPC2DPGRYYtQJie1/AbKn7rix3XxC2bPySSpe3/3pa1NSQY3JfLZPfhTvOv5X3asU3d+CN8cHFNly21HJCqdpPYYXJnZ0Q248KbWxdJSYkpPA+PVGLoa+VPrHjEHhMMb9gvl1ynwuRV4rheW49iNHCcsgxAu18=",
	SdkSign: "fwU0VSlfsV+RE6SccuAqu/fZItgFDmxD6kCwcOrchbfyT8SiY7w60HzYWmEU 9EWcGO+a3s9ru+AePT2JeBS3ml0Pk6k4RM1PrdUJxUvnie+Opf+u6ygrTDHT W6R2c3XVSfROaT0hqXFxPI82XRCUtwr61RhUtBxG2hZiF3sjU2URVJdCVuvr gw==",
	YwKey:   "ywuldVe83KrJ",
	YwGuid:  "525165249",
	//仅签到 新用户 没法看视频
	TaskType: []TaskType{},
}
var configs = []QiDianApiConfig{
	QiDianApiConfig{
		QdInfo:  "SO+aPyWTJ02k4C9FkkB29fACDXIsJx4pAGbhVI07D8hjHPOEsCFgpJ99gS3kYIjunO+UrcWbhPgIlUSo3XxdoisFnouWF80qfP+9nYAPZWuWE/x7ukJhxmRzdw1aVsBmIiveq7vRwg3TXyIs43bQR3QvtRub8keYvoqD6sPrygjYPGJO9epFkB+fl7WpyOWndaOEzW8rBQsZsBlYOgisDB2wfkLKY4MRmjXtJXExeryCF60qFFMggzPdI0Ix+8Bd4wq2H4FjKDpW/8bUd251Pcqp/aFpa6+ta4mpDrYTvSyPQRQ4L4c3TQ==",
		SdkSign: "fwU0VSlfsV+RE6SccuAquwkzDajTNoEiqiGCwQfq7fHo7uruH5m4d6dpHZ7l UMFsNXq7acQdUlrZaOxwvjFn8/1Q8jvP0T+BNavJZeopiM22wFQTe3MJWYcb Qa8F7Za38bJHVI9c24MBjk26vMafbmqA6dobIfjQHDPsPGzpoK9DW3wwCs9H WG7t8CVhvVUpWLbh3WXeFco=",
		YwKey:   "ywv57qBHtL3q",
		YwGuid:  "460067960",
		//不玩游戏 目前没有破解imei相关的加密,多账号时只能使用多台手机的cookie,否则后台会不计算时长
		TaskType: []TaskType{TPSurpriseBenefit, TPDailyBenefit, TPVideoRewardTabTaskList},
	},
	QiDianApiConfig{
		QdInfo:  "SO+aPyWTJ02k4C9FkkB29fACDXIsJx4pAGbhVI07D8hjHPOEsCFgpJ99gS3kYIjunO+UrcWbhPgIlUSo3XxdoisFnouWF80qfP+9nYAPZWuWE/x7ukJhxmRzdw1aVsBmIiveq7vRwg0jmRKWIbLk0ya0DYM4609KsFwEIC7L1WVJQXRx5UNlnsjgPb3hCSa7HUCpa9zqtpcyDyM2fWAayh4ikkxUyKXzIv/ZJHgdpi96D9QKK7YGVyioz85qI/w483DdS4qJgDVUP2YdEyq0NAT7LIZm92+jzv79Zpu9Q5k+xJ/ThJI6CQ==",
		SdkSign: "fwU0VSlfsV/z4GQNcgcrcGiuXtCeOrancxEZUexJzB4Bt5Ne33V01rzTHpj0 sN7tTN1GDScbR3MXVE1RmmcnSmV8CEBCdSKOiWi7OWmLtqt1RzTPszFLWG/U xLK5yCM7lNhPLv/y6CekqR5JAXmNLkSnyvdfTxDk+otpHi+PMEgC/U1IfwUj OzH3JdbgQP7OKXwpvUozje0=",
		YwKey:   "ywqpRoQfqBKt",
		YwGuid:  "460067960",
		//不玩游戏 目前没有破解imei相关的加密,多账号时只能使用多台手机的cookie,否则后台会不计算时长
		TaskType: []TaskType{TPSurpriseBenefit, TPDailyBenefit, TPVideoRewardTabTaskList},
	},
}

func getApi(config QiDianApiConfig) *QiDianApi {
	meta, err := sign.NewMeta(config.QdInfo, config.SdkSign)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	ret := NewQiDianApi(meta, config.YwKey, config.YwGuid)
	ret.Cli = GetProxyClient()
	return ret
}
func init() {
	api = getApi(configs[0])
	api2 := getApi(configs[1])
	fmt.Printf("%v\n", api2.sign)
}
func TestCheckIn(t *testing.T) {
	if resp, err := api.CheckIn(); err != nil {
		t.Error(err)
	} else {
		t.Log(resp)
	}
}
func TestAdvMainPage(t *testing.T) {
	if resp, err := api.AdvMainPage(); err != nil {
		t.Error(err)
	} else {
		t.Log(resp)
	}
}
func TestFinishWatch_SurpriseBenefit(t *testing.T) {
	advMainPage, err := api.AdvMainPage()
	if err != nil {
		t.Error(err)
	}
	surprise := advMainPage.GetSurpriseBenefit()
	if surprise.IsFinished == 0 {
		if resp, err := api.FinishWatch(surprise.TaskId); err != nil {
			t.Error(err)
		} else {
			t.Log(resp)
		}
	}
}

func heartBeat(api *QiDianApi) {
	for i := 0; i < 10; i++ {
		beat, err := api.UrlHeartBeat()
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		fmt.Printf("%v\n", beat)
		time.Sleep(time.Second * 30)
	}
}
func TestHeartBeat(t *testing.T) {
	api2 := getApi(configs[0])
	api1 := getApi(configs[1])
	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		defer wg.Done()
		heartBeat(api1)
	}()
	time.Sleep(time.Second * 15)
	heartBeat(api2)
	wg.Wait()
}
func TestFinishWatch_DailyBenefit(t *testing.T) {
	advMainPage, err := api.AdvMainPage()
	if err != nil {
		t.Error(err)
	}
	list := advMainPage.GetDailyBenefitTaskList()
	for _, task := range list {
		if task.IsFinished == 0 {
			if resp, err := api.FinishWatch(task.TaskId); err != nil {
				t.Error(err)
			} else {
				t.Log(resp)
			}
			time.Sleep(time.Second * 25)
		}
	}

}

func TestFinishWatch_VideoRewardTabTaskList(t *testing.T) {
	advMainPage, err := api.AdvMainPage()
	if err != nil {
		t.Error(err)
	}
	list := advMainPage.GetVideoRewardTabTaskList(true)
	for _, task := range list {
		if task.IsFinished == 0 {
			for i := task.Process; i < task.Total; i++ {
				if resp, err := api.FinishWatch(task.TaskId); err != nil {
					t.Error(err)
				} else {
					t.Log(resp)
				}
				time.Sleep(time.Second * 25)
			}

		}
	}
}
func TestFinishWatch_MoreRewardTabTaskList(t *testing.T) {
	advMainPage, err := api.AdvMainPage()
	if err != nil {
		t.Error(err)
	}
	list := advMainPage.GetMoreRewardTabTaskList()
	for _, task := range list {
		if task.IsFinished == 0 {
			if resp, err := api.ReceiveTaskReward(task.TaskId); err != nil {
				t.Error(err)
			} else {
				t.Log(resp)
			}
			time.Sleep(time.Second)
		}
	}
}
