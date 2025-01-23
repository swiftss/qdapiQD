package qdapi

import (
	"fmt"
	"github.com/pzx521521/qdapi/sign"
	"testing"
	"time"
)

var api *QiDianApi

var c = QiDianApiConfig{
	QDInfo:  "SO+aPyWTJ02k4C9FkkB29fACDXIsJx4pAGbhVI07D8hjHPOEsCFgpJ99gS3kYIjunO+UrcWbhPgIlUSo3XxdoisFnouWF80qfP+9nYAPZWuWE/x7ukJhxmRzdw1aVsBmIiveq7vRwg2cWq9rUSiFr8+TzhhMqzsahrzx8MySAqCgRREoe3UHQeRJdOptPZy7Zp80Got/jx81dN6SGzvP1ghm8ebYnQ4uKmeMn1XaGn/TLWTSWjEORIcCumvlLP9QfH1MAE6JFjwDBFxKiHmJ8gEIH575OkGuudZIY/axUcj4jDcvhOR5Tw==",
	SDKSign: "fwU0VSlfsV/NtCFBjpJaraJc+v78FE6Ksfwf1tERIfrtPu9CbZQRyd86GyeD IClued440A6dZPtUTCEm/Nmg/sJUmk6piMqL/1oIUjvvyDgxV8JLiODJGWWj nE13omPAbgjI/g9dIOH6GIHl2Kqs8NCcyMpf4AsgXg2+qku6oG9QCbptHafw zhcRB8rTY0M5BqQnBW7JA2I=",
	YWKey:   "ywROdPzlJ8Tp",
	YWGuid:  "460067960",
}

func init() {
	meta, err := sign.NewMeta(c.QDInfo, c.SDKSign)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	api = NewQiDianApi(meta, c.YWKey, c.YWGuid)
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

func TestHeartBeat(t *testing.T) {
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
