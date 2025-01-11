package main

import (
	"fmt"
	"qidian/sign"
)

const PhoneQDInfo = "SO+aPyWTJ02k4C9FkkB29fACDXIsJx4pAGbhVI07D8hjHPOEsCFgpJ99gS3kYIjunO+UrcWbhPgIlUSo3XxdoisFnouWF80qfP+9nYAPZWviV9DlRVtRllf2xwn3SDILgUVhzxJXrTRNcGeynaP07zVZ5qe7MsQgKlxQdWxM6mFdhlYjvxrV+vON3pGlpR6i99QzTXesSmhLrNUXyxfExycfosXSayIx7cg++mgVmMuYlzq0lHdkLtE9Xy/osz0yxcsC+f+qlmWM7h/koIE014cWbRBsHjwBQJJkU6h2fa5lDRNhYZLnfQ=="
const PhoneSDKSign = "fwU0VSlfsV/NtCFBjpJarbYpi9mlbLU/EDzhOVoz2RdtheX+SLpjTy8L2+gA InschgJSs1O5vbtFpSZ6+GPI8iEd6QhtwlTz8ODKLNM1r+aH0A8sY5+lP6la DPt/GpDgPvW5ZvKHiqnIqFEJHRoPYEshR2+cAq03JfcYLPvSfE7DpuHLVA2F mRtLGCdVWmTujc/5Lb+/Cmk="

const YWKEY = "ywU8TfcHg8J4"
const YWGUID = "120154865151"

func main() {
	meta, err := sign.NewMeta(PhoneQDInfo, PhoneSDKSign)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	api := NewQiDianApi(meta, YWKEY, YWGUID)
	//for charles
	//api.Cli = GetProxyClient()
	//for github action
	api.Cli = GetInsecureClient()
	DoTask(api,
		//一小时一个的宝箱
		TPSurpriseBenefit,
		//每天的8个任务
		TPDailyBenefit,
		//看3个得10点的任务
		TPVideoRewardTabTaskList,
		//更多任务 游戏+30点 等等  暂不支持
		TPMoreRewardTab)
}
