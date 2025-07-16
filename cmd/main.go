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
		QdInfo:   "SO+aPyWTJ02k4C9FkkB29fACDXIsJx4pAGbhVI07D8hjHPOEsCFgpKs6sLj9MLaOMad0nLWkjLZiPGzrYzIs3aToDVNBVw9rru4owdN+fTUJ1mfNNfX+Cf/Jmpzp3FOyeRsYzY8OzU10SPvkhqZrJ7RYZvpNqrFJMRQhrfuh+MYT4aZ8vxffAtAoPRlQoRBy5xSAFWjk66R+eBH6Vr4SZd4WvIJhjPa4TDOD4O3M0OWUXmoRQgSyvdN3y8apTukrM0zmAgzSKoxqNPb7x2bzM3X1+pazZ+RZBofH0JUhsNkMLZ6M6QQIyQ==",
		SdkSign:  "fwU0VSlfsV+u7EEafvZ0VP/62Lz482/spUt+fYCLtGsY/kr5NevIloLswUFz MMgxHqh9znYMH7Sc0nTYj0dblgVvuhZSCg/tRXIBRAIxVYYVp7ng0MfabciM QH9MKiOzA6h1DJqW8YigTPfBGyIQL96eAATjsngLinxlm6bWLWRPkc43whK6 xAz8o/FpXnR+yL9+DRaFYVU=",
		YwKey:    "ykGkxcJ86KGa",
		YwGuid:   "120154865151",
		TaskType: []qdapi.TaskType{qdapi.TPSurpriseBenefit, qdapi.TPDailyBenefit, qdapi.TPVideoRewardTabTaskList, qdapi.TPMoreRewardTab},
	},
	qdapi.QiDianApiConfig{
		QdInfo:  "SO+aPyWTJ02k4C9FkkB29fACDXIsJx4pAGbhVI07D8hjHPOEsCFgpJ99gS3kYIjunO+UrcWbhPgIlUSo3XxdoisFnouWF80qfP+9nYAPZWuWE/x7ukJhxmRzdw1aVsBmIiveq7vRwg0jmRKWIbLk0ya0DYM4609KsFwEIC7L1WVJQXRx5UNlnsjgPb3hCSa7HUCpa9zqtpcyDyM2fWAayh4ikkxUyKXzIv",
		SdkSign: "fwU0VSlfsV+u7EEafvZ0VJ730Gueop3ofhHZvpfPxzhpOECAeBYcgqSo01uN hDc5gW3eQm4IzEDPHcky6Ocd7GcYRRyYbV64JUNYUGSfgkTzh1Fe36KG4+DR ISl1s9ovCYcIt59SkB04vvZntgnIXn1ptEsWBJbw9uRY3yZYY3VLBV5m7kcv iibvLJU4Txbry2iGc24ll9g=",
		YwKey:   "ykdAwoTvGwJ9",
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
