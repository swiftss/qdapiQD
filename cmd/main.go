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
		YwKey:    "yw0m8sFvXRqa",
		YwGuid:   "120154865151",
		TaskType: []qdapi.TaskType{qdapi.TPSurpriseBenefit, qdapi.TPDailyBenefit, qdapi.TPVideoRewardTabTaskList, qdapi.TPMoreRewardTab},
	},
	qdapi.QiDianApiConfig{
		QdInfo:  "SO+aPyWTJ02k4C9FkkB29fACDXIsJx4pAGbhVI07D8hjHPOEsCFgpJ99gS3kYIjunO+UrcWbhPgIlUSo3XxdoisFnouWF80qfP+9nYAPZWuWE/x7ukJhxmRzdw1aVsBmIiveq7vRwg0jmRKWIbLk0ya0DYM4609KsFwEIC7L1WVJQXRx5UNlnsjgPb3hCSa7HUCpa9zqtpcyDyM2fWAayh4ikkxUyKXzIv/ZJHgdpi96D9QKK7YGVyioz85qI/w483DdS4qJgDVUP2YdEyq0NAT7LIZm92+jzv79Zpu9Q5k+xJ/ThJI6CQ==",
		SdkSign: "fwU0VSlfsV/z4GQNcgcrcGiuXtCeOrancxEZUexJzB4Bt5Ne33V01rzTHpj0 sN7tTN1GDScbR3MXVE1RmmcnSmV8CEBCdSKOiWi7OWmLtqt1RzTPszFLWG/U xLK5yCM7lNhPLv/y6CekqR5JAXmNLkSnyvdfTxDk+otpHi+PMEgC/U1IfwUj OzH3JdbgQP7OKXwpvUozje0=",
		YwKey:   "ywEw0mB2ErUg",
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
