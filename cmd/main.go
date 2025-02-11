package main

import (
	"fmt"
	"github.com/pzx521521/qdapi"
	"github.com/pzx521521/qdapi/sign"
	"log"
	"net/http"
	"runtime"
	"sync"
)

var configs = []qdapi.QiDianApiConfig{
	qdapi.QiDianApiConfig{
		QdInfo:   "SO+aPyWTJ02k4C9FkkB29fACDXIsJx4pAGbhVI07D8hjHPOEsCFgpJ99gS3kYIjunO+UrcWbhPgIlUSo3XxdoisFnouWF80qfP+9nYAPZWuWE/x7ukJhxq8DEJW+n90UAoC6t3e9KFYiTLEscTjh73LcOjrTQz+qqJjYEZbh6kASXSCx91sFGt5BnhiBB/3z4VWhKNmFjaF4Kxw7l6OL/MPyJ5eI4116Y43wcvg2k1TjTr0n/LqBAUl0hz82734VOvc95DeiNUDQKneKfLUDPeARmS5sYPNGRFNFr36HmPTkbfd+w6OFVA==",
		SdkSign:  "fwU0VSlfsV8VwoTAKeiLPkPXuEgmx4T3/qIzsXOXgNt+iZfAcWS+jVkp1l8o 5kOozMtfRIgSXO5lM3InyuNrlINApdhLJlFnaYIchu5wBgNoqSc5fPuEJDxd Sdp1VTvMsqmfmjwCHO5Iu2NpGJmNOFF6ilOIS67pU1ESOqJz+THNQFNCxGmn a6CsLsOPn0Gyv7Jzg2Rvxag=",
		YwKey:    "yw39rD5TJ4TK",
		YwGuid:   "120154865151",
		TaskType: []qdapi.TaskType{qdapi.TPSurpriseBenefit, qdapi.TPDailyBenefit, qdapi.TPVideoRewardTabTaskList, qdapi.TPMoreRewardTab},
	},
	qdapi.QiDianApiConfig{
		QdInfo:  "WnYe23HKHHbFexRmDMhQphCTehSF/SNEL1SCOpnokmfIUobvKBUfPf9v9OO9vQkVpLUc2ZGEfF+RYPPC2DPGRYYtQJie1/AbKn7rix3XxC2bPySSpe3/3pa1NSQY3JfLZPfhTvOv5X3asU3d+CN8cHFNly21HJCqdpPYYXJnZ0Q248KbWxdJSYkpPA+PVGLoa+VPrHjEHhMMb9gvl1ynwuRV4rheW49iNHCcsgxAu18=",
		SdkSign: "fwU0VSlfsV+RE6SccuAqu/fZItgFDmxD6kCwcOrchbfyT8SiY7w60HzYWmEU 9EWcGO+a3s9ru+AePT2JeBS3ml0Pk6k4RM1PrdUJxUvnie+Opf+u6ygrTDHT W6R2c3XVSfROaT0hqXFxPI82XRCUtwr61RhUtBxG2hZiF3sjU2URVJdCVuvr gw==",
		YwKey:   "ywuldVe83KrJ",
		YwGuid:  "525165249",
		//仅签到 新用户 没法看视频
		TaskType: []qdapi.TaskType{},
	},
	qdapi.QiDianApiConfig{
		QdInfo:  "SO+aPyWTJ02k4C9FkkB29fACDXIsJx4pAGbhVI07D8hjHPOEsCFgpJ99gS3kYIjunO+UrcWbhPgIlUSo3XxdoisFnouWF80qfP+9nYAPZWuWE/x7ukJhxmRzdw1aVsBmIiveq7vRwg3TXyIs43bQR3QvtRub8keYvoqD6sPrygjYPGJO9epFkB+fl7WpyOWndaOEzW8rBQsZsBlYOgisDB2wfkLKY4MRmjXtJXExeryCF60qFFMggzPdI0Ix+8Bd4wq2H4FjKDpW/8bUd251Pcqp/aFpa6+ta4mpDrYTvSyPQRQ4L4c3TQ==",
		SdkSign: "fwU0VSlfsV+RE6SccuAquwkzDajTNoEiqiGCwQfq7fHo7uruH5m4d6dpHZ7l UMFsNXq7acQdUlrZaOxwvjFn8/1Q8jvP0T+BNavJZeopiM22wFQTe3MJWYcb Qa8F7Za38bJHVI9c24MBjk26vMafbmqA6dobIfjQHDPsPGzpoK9DW3wwCs9H WG7t8CVhvVUpWLbh3WXeFco=",
		YwKey:   "ywv57qBHtL3q",
		YwGuid:  "460067960",
		//不玩游戏 目前没有破解imei相关的加密,多账号时只能使用多台手机的cookie,否则后台会不计算时长
		TaskType: []qdapi.TaskType{qdapi.TPSurpriseBenefit, qdapi.TPDailyBenefit, qdapi.TPVideoRewardTabTaskList},
	},
}

func main() {
	var cli *http.Client
	if runtime.GOOS == "darwin" {
		//for charles
		cli = qdapi.GetProxyClient()
	} else {
		//for github action
		cli = qdapi.GetInsecureClient()
	}
	CheckInAndDoTaskMulti(cli, configs...)
}
func CheckInAndDoTaskMulti(cli *http.Client, configs ...qdapi.QiDianApiConfig) {
	var wg sync.WaitGroup
	for i, config := range configs {
		wg.Add(1)
		go func(index int, config qdapi.QiDianApiConfig) {
			CheckInAndDoTask(cli, config)
			wg.Done()
		}(i, config)
	}
	wg.Wait()
}

func CheckInAndDoTask(client *http.Client, config qdapi.QiDianApiConfig) {
	meta, err := sign.NewMeta(config.QdInfo, config.SdkSign)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}

	log.Printf("%v\n", meta)
	api := qdapi.NewQiDianApi(meta, config.YwKey, config.YwGuid)
	api.Cli = client
	resp, err := api.CheckIn()
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	log.Printf("%s:%v\n", api.TipName(), resp)
	err = qdapi.DoTask(api, config.TaskType...)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}
