package main

import (
	"fmt"
	"github.com/pzx521521/qdapi"
	"github.com/pzx521521/qdapi/sign"
	"log"
	"sync"
)

var configs = []qdapi.QiDianApiConfig{
	qdapi.QiDianApiConfig{
		QDInfo:  "SO+aPyWTJ02k4C9FkkB29fACDXIsJx4pAGbhVI07D8hjHPOEsCFgpJ99gS3kYIjunO+UrcWbhPgIlUSo3XxdoisFnouWF80qfP+9nYAPZWviV9DlRVtRllf2xwn3SDILgUVhzxJXrTRNcGeynaP07zVZ5qe7MsQgKlxQdWxM6mFdhlYjvxrV+vON3pGlpR6i99QzTXesSmhLrNUXyxfExycfosXSayIx7cg++mgVmMuYlzq0lHdkLtE9Xy/osz0yxcsC+f+qlmWM7h/koIE014cWbRBsHjwBQJJkU6h2fa5lDRNhYZLnfQ==",
		SDKSign: "fwU0VSlfsV/NtCFBjpJarbYpi9mlbLU/EDzhOVoz2RdtheX+SLpjTy8L2+gA InschgJSs1O5vbtFpSZ6+GPI8iEd6QhtwlTz8ODKLNM1r+aH0A8sY5+lP6la DPt/GpDgPvW5ZvKHiqnIqFEJHRoPYEshR2+cAq03JfcYLPvSfE7DpuHLVA2F mRtLGCdVWmTujc/5Lb+/Cmk=",
		YWKey:   "ywU8TfcHg8J4",
		YWGuid:  "120154865151",
	},
	qdapi.QiDianApiConfig{
		QDInfo:  "SO+aPyWTJ02k4C9FkkB29fACDXIsJx4pAGbhVI07D8hjHPOEsCFgpJ99gS3kYIjunO+UrcWbhPgIlUSo3XxdoisFnouWF80qfP+9nYAPZWuWE/x7ukJhxmRzdw1aVsBmIiveq7vRwg2cWq9rUSiFr8+TzhhMqzsahrzx8MySAqCgRREoe3UHQeRJdOptPZy7Zp80Got/jx81dN6SGzvP1ghm8ebYnQ4uKmeMn1XaGn/TLWTSWjEORIcCumvlLP9QfH1MAE6JFjwDBFxKiHmJ8gEIH575OkGuudZIY/axUcj4jDcvhOR5Tw==",
		SDKSign: "fwU0VSlfsV/NtCFBjpJaraJc+v78FE6Ksfwf1tERIfrtPu9CbZQRyd86GyeD IClued440A6dZPtUTCEm/Nmg/sJUmk6piMqL/1oIUjvvyDgxV8JLiODJGWWj nE13omPAbgjI/g9dIOH6GIHl2Kqs8NCcyMpf4AsgXg2+qku6oG9QCbptHafw zhcRB8rTY0M5BqQnBW7JA2I=",
		YWKey:   "ywROdPzlJ8Tp",
		YWGuid:  "460067960",
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
			CheckInAndDoTask(config.QDInfo, config.SDKSign, config.YWKey, config.YWGuid)
			wg.Done()
		}(config)
	}
	wg.Wait()
}

func CheckInAndDoTask(qDInfo, sDKSign, ywKey, ywGuid string) {
	meta, err := sign.NewMeta(qDInfo, sDKSign)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	log.Printf("%v\n", meta)
	api := qdapi.NewQiDianApi(meta, ywKey, ywGuid)
	//for charles
	//api.Cli = qdapi.GetProxyClient()
	//for github action
	api.Cli = qdapi.GetInsecureClient()
	resp, err := api.CheckIn()
	if err != nil {
		return
	}
	log.Printf("%v\n", resp)
	err = qdapi.DoTask(api,
		//一小时一个的宝箱
		qdapi.TPSurpriseBenefit,
		//每天的8个任务
		qdapi.TPDailyBenefit,
		//看3个得10点的任务
		qdapi.TPVideoRewardTabTaskList,
		//更多任务 游戏+30点 等等  暂不支持
		qdapi.TPMoreRewardTab)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}
