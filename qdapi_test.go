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
		QdInfo:   "SO+aPyWTJ02k4C9FkkB29fACDXIsJx4pAGbhVI07D8hjHPOEsCFgpKs6sLj9MLaOMad0nLWkjLZiPGzrYzIs3aToDVNBVw9rru4owdN+fTUJ1mfNNfX+Cf/Jmpzp3FOyeRsYzY8OzU10SPvkhqZrJ7RYZvpNqrFJMRQhrfuh+MYT4aZ8vxffAtAoPRlQoRBy5xSAFWjk66R+eBH6Vr4SZd4WvIJhjPa4TDOD4O3M0OWUXmoRQgSyvdN3y8apTukrM0zmAgzSKoxqNPb7x2bzM3X1+pazZ+RZBofH0JUhsNkMLZ6M6QQIyQ==",
		SdkSign:  "fwU0VSlfsV+u7EEafvZ0VP/62Lz482/spUt+fYCLtGsY/kr5NevIloLswUFz MMgxHqh9znYMH7Sc0nTYj0dblgVvuhZSCg/tRXIBRAIxVYYVp7ng0MfabciM QH9MKiOzA6h1DJqW8YigTPfBGyIQL96eAATjsngLinxlm6bWLWRPkc43whK6 xAz8o/FpXnR+yL9+DRaFYVU=",
		YwKey:    "ykGkxcJ86KGa",
		YwGuid:   "120154865151",
		TaskType: []TaskType{TPSurpriseBenefit, TPDailyBenefit, TPVideoRewardTabTaskList, TPMoreRewardTab},
	},
	QiDianApiConfig{
		QdInfo:  "SO+aPyWTJ02k4C9FkkB29fACDXIsJx4pAGbhVI07D8hjHPOEsCFgpKs6sLj9MLaOMad0nLWkjLZiPGzrYzIs3aToDVNBVw9rru4owdN+fTUJ1mfNNfX+CQBZaD0PyG/b1t2Pnsyhhg3pJ2nBg9B2wZLflsDrxvbyZS0ph0Zady3Fdfzhtv74uNSvheBPj8pQNMc4xpwZiCHd/CUPnt+qgDIobBzd6DgRQ+BEd7LZuRRwl5p2W+FOeBhW/Bh08BLoBopv/siCjM8FZhGOlZRC929oO/oLCyCh0ok6hmVXVXNHbXI/b27nbg==",
		SdkSign: "fwU0VSlfsV+u7EEafvZ0VJ730Gueop3oqm+NZKzff7PV69FBxLEOduSkSAG1 xp/AXsOxQQOn4cKzhbpCf+qs1IQ8QtJlB3qKQwxM81DejQqjedkmm7CUpxZT klBDkdwTDw0iGh63ZivS0htCBXvWyS21r63gm6xMyfzw/pfc3kVp/OnvKA6M CSuOlMJiEXwgOk2nJQXiW34=",
		YwKey:   "ykdAwoTvGwJ9",
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

func TestConfig(t *testing.T) {
	err := SaveConfigToJSON("./config.json", configs)
	if err != nil {
		return
	}
	json, err := LoadConfigFromJSON("./config.json")
	if err != nil {
		return
	}
	fmt.Printf("%v\n", json)
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
