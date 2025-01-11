package qdapi

import (
	"fmt"
	"github.com/pzx521521/qdapi/sign"
	"testing"
	"time"
)

var api *QiDianApi

const PhoneQDInfo = "SO+aPyWTJ02k4C9FkkB29fACDXIsJx4pAGbhVI07D8hjHPOEsCFgpJ99gS3kYIjunO+UrcWbhPgIlUSo3XxdoisFnouWF80qfP+9nYAPZWviV9DlRVtRllf2xwn3SDILgUVhzxJXrTRNcGeynaP07zVZ5qe7MsQgKlxQdWxM6mFdhlYjvxrV+vON3pGlpR6i99QzTXesSmhLrNUXyxfExycfosXSayIx7cg++mgVmMuYlzq0lHdkLtE9Xy/osz0yxcsC+f+qlmWM7h/koIE014cWbRBsHjwBQJJkU6h2fa5lDRNhYZLnfQ=="
const PhoneSDKSign = "fwU0VSlfsV/NtCFBjpJarbYpi9mlbLU/EDzhOVoz2RdtheX+SLpjTy8L2+gA InschgJSs1O5vbtFpSZ6+GPI8iEd6QhtwlTz8ODKLNM1r+aH0A8sY5+lP6la DPt/GpDgPvW5ZvKHiqnIqFEJHRoPYEshR2+cAq03JfcYLPvSfE7DpuHLVA2F mRtLGCdVWmTujc/5Lb+/Cmk="

const YWKEY = ""
const YWGUID = ""

func init() {
	meta, err := sign.NewMeta(PhoneQDInfo, PhoneSDKSign)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	api = NewQiDianApi(meta, YWKEY, YWGUID)
	api.Cli = GetProxyClient()
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
			if resp, err := api.FinishWatch(task.TaskId); err != nil {
				t.Error(err)
			} else {
				t.Log(resp)
			}
			time.Sleep(time.Second * 25)
		}
	}
}
