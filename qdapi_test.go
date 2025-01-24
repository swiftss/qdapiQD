package qdapi

import (
	"fmt"
	"github.com/pzx521521/qdapi/sign"
	"testing"
	"time"
)

var api *QiDianApi

var c = QiDianApiConfig{
	QDInfo:  "WnYe23HKHHbFexRmDMhQphCTehSF/SNEL1SCOpnokmfIUobvKBUfPf9v9OO9vQkVpLUc2ZGEfF+RYPPC2DPGRYYtQJie1/AbKn7rix3XxC2bPySSpe3/3pa1NSQY3JfLZPfhTvOv5X3asU3d+CN8cHFNly21HJCqdpPYYXJnZ0Q248KbWxdJSYkpPA+PVGLoa+VPrHjEHhMMb9gvl1ynwuRV4rheW49iNHCcsgxAu18=",
	SDKSign: "fwU0VSlfsV+RE6SccuAqu/fZItgFDmxD6kCwcOrchbfyT8SiY7w60HzYWmEU 9EWcGO+a3s9ru+AePT2JeBS3ml0Pk6k4RM1PrdUJxUvnie+Opf+u6ygrTDHT W6R2c3XVSfROaT0hqXFxPI82XRCUtwr61RhUtBxG2hZiF3sjU2URVJdCVuvr gw==",
	YWKey:   "ywuldVe83KrJ",
	YWGuid:  "525165249",
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
