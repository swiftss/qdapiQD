package main

import (
	"fmt"
	"github.com/pzx521521/qdapi"
	"github.com/pzx521521/qdapi/sign"
	"log"
	"runtime"
	"sync"
)

var configs = []qdapi.QiDianApiConfig{
	qdapi.QiDianApiConfig{
		QDInfo:   "SO+aPyWTJ02k4C9FkkB29fACDXIsJx4pAGbhVI07D8hjHPOEsCFgpJ99gS3kYIjunO+UrcWbhPgIlUSo3XxdoisFnouWF80qfP+9nYAPZWviV9DlRVtRllf2xwn3SDILgUVhzxJXrTRNcGeynaP07zVZ5qe7MsQgKlxQdWxM6mFdhlYjvxrV+vON3pGlpR6i99QzTXesSmhLrNUXyxfExycfosXSayIx7cg++mgVmMuYlzq0lHdkLtE9Xy/osz0yxcsC+f+qlmWM7h/koIE014cWbRBsHjwBQJJkU6h2fa5lDRNhYZLnfQ==",
		SDKSign:  "fwU0VSlfsV/NtCFBjpJarbYpi9mlbLU/EDzhOVoz2RdtheX+SLpjTy8L2+gA InschgJSs1O5vbtFpSZ6+GPI8iEd6QhtwlTz8ODKLNM1r+aH0A8sY5+lP6la DPt/GpDgPvW5ZvKHiqnIqFEJHRoPYEshR2+cAq03JfcYLPvSfE7DpuHLVA2F mRtLGCdVWmTujc/5Lb+/Cmk=",
		YWKey:    "ywU8TfcHg8J4",
		YWGuid:   "120154865151",
		TaskType: []qdapi.TaskType{qdapi.TPSurpriseBenefit, qdapi.TPDailyBenefit, qdapi.TPVideoRewardTabTaskList, qdapi.TPMoreRewardTab},
	},
	qdapi.QiDianApiConfig{
		QDInfo:   "WnYe23HKHHbFexRmDMhQphCTehSF/SNEL1SCOpnokmfIUobvKBUfPf9v9OO9vQkVpLUc2ZGEfF+RYPPC2DPGRYYtQJie1/AbKn7rix3XxC2bPySSpe3/3pa1NSQY3JfLZPfhTvOv5X3asU3d+CN8cHFNly21HJCqdpPYYXJnZ0Q248KbWxdJSYkpPA+PVGLoa+VPrHjEHhMMb9gvl1ynwuRV4rheW49iNHCcsgxAu18=",
		SDKSign:  "fwU0VSlfsV+RE6SccuAqu/fZItgFDmxD6kCwcOrchbfyT8SiY7w60HzYWmEU 9EWcGO+a3s9ru+AePT2JeBS3ml0Pk6k4RM1PrdUJxUvnie+Opf+u6ygrTDHT W6R2c3XVSfROaT0hqXFxPI82XRCUtwr61RhUtBxG2hZiF3sjU2URVJdCVuvr gw==",
		YWKey:    "ywuldVe83KrJ",
		YWGuid:   "525165249",
		TaskType: []qdapi.TaskType{},
	},
	qdapi.QiDianApiConfig{
		QDInfo:   "SO+aPyWTJ02k4C9FkkB29fACDXIsJx4pAGbhVI07D8hjHPOEsCFgpJ99gS3kYIjunO+UrcWbhPgIlUSo3XxdoisFnouWF80qfP+9nYAPZWuWE/x7ukJhxmRzdw1aVsBmIiveq7vRwg3TXyIs43bQR3QvtRub8keYvoqD6sPrygjYPGJO9epFkB+fl7WpyOWndaOEzW8rBQsZsBlYOgisDB2wfkLKY4MRmjXtJXExeryCF60qFFMggzPdI0Ix+8Bd4wq2H4FjKDpW/8bUd251Pcqp/aFpa6+ta4mpDrYTvSyPQRQ4L4c3TQ==",
		SDKSign:  "fwU0VSlfsV+RE6SccuAquwkzDajTNoEiqiGCwQfq7fHo7uruH5m4d6dpHZ7l UMFsNXq7acQdUlrZaOxwvjFn8/1Q8jvP0T+BNavJZeopiM22wFQTe3MJWYcb Qa8F7Za38bJHVI9c24MBjk26vMafbmqA6dobIfjQHDPsPGzpoK9DW3wwCs9H WG7t8CVhvVUpWLbh3WXeFco=",
		YWKey:    "ywv57qBHtL3q",
		YWGuid:   "460067960",
		TaskType: []qdapi.TaskType{qdapi.TPSurpriseBenefit, qdapi.TPDailyBenefit, qdapi.TPVideoRewardTabTaskList, qdapi.TPMoreRewardTab},
	},
}

func main() {
	CheckInAndDoTaskMulti(configs...)
}
func CheckInAndDoTaskMulti(configs ...qdapi.QiDianApiConfig) {
	var wg sync.WaitGroup
	for _, config := range configs {
		wg.Add(1)
		go func(config qdapi.QiDianApiConfig) {
			CheckInAndDoTask(config)
			wg.Done()
		}(config)
	}
	wg.Wait()
}

func CheckInAndDoTask(config qdapi.QiDianApiConfig) {
	meta, err := sign.NewMeta(config.QDInfo, config.SDKSign)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}

	log.Printf("%v\n", meta)
	api := qdapi.NewQiDianApi(meta, config.YWKey, config.YWGuid)

	if runtime.GOOS == "darwin" {
		//for charles
		api.Cli = qdapi.GetProxyClient()
	} else {
		//for github action
		api.Cli = qdapi.GetInsecureClient()
	}
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
