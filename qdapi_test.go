package main

import (
	"fmt"
	"qidian/sign"
	"testing"
	"time"
)

var api *QiDianApi

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
