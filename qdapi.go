package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"qidian/sign"
)

func Request[T any](qd *QiDianApi, uri string, method string, data interface{}) (*T, error) {
	get, err := qd.request(uri, method, data)
	if err != nil {
		return nil, err
	}
	var ret *T
	err = json.Unmarshal(get, &ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}

type QiDianApi struct {
	sign   *sign.Meta
	Cli    *http.Client
	ywkey  string
	ywguid string
}

func NewQiDianApi(sign *sign.Meta, ywkey, ywguid string) *QiDianApi {
	return &QiDianApi{
		sign:   sign,
		Cli:    http.DefaultClient,
		ywkey:  ywkey,
		ywguid: ywguid,
	}
}
func (qd *QiDianApi) CheckIn() (*CheckinResp, error) {
	return Request[CheckinResp](qd, UrlCheckIn, http.MethodPost, nil)
}

func (qd *QiDianApi) AdvMainPage() (*AdvMainPage, error) {
	return Request[AdvMainPage](qd, UrlAdvMainPage, http.MethodGet, nil)
}
func (qd *QiDianApi) FinishWatch(Id string) (*FinishWatch, error) {
	data := "taskId=" + Id + "&BanId=0&BanMessage=&CaptchaAId=&CaptchaType=0&CaptchaURL=&Challenge=&Gt=&NewCaptcha=0&Offline=0&PhoneNumber=&SessionKey="
	return Request[FinishWatch](qd, UrlFinishWatch, http.MethodPost, data)
}

func (qd *QiDianApi) header(req *http.Request) error {
	sdkSign, err := qd.sign.SDKSign(req.URL.String())
	if err != nil {
		return err
	}
	req.Header.Set("SDKSign", sdkSign)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 14; 2106118C Build/UKQ1.231207.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/129.0.6668.100 Mobile Safari/537.36")
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 14; 2106118C Build/UKQ1.231207.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/129.0.6668.100 Mobile Safari/537.36 QDJSSDK/1.0  QDNightStyle_1  QDReaderAndroid/7.9.384/1466/1002140/Xiaomi/QDShowNativeLoading")
	return nil
}
func (qd *QiDianApi) request(uri string, method string, data any) ([]byte, error) {
	err := qd.sign.ModifyTimeStamp()
	if err != nil {
		return nil, err
	}
	qDInfo, err := qd.sign.QDInfo()
	if err != nil {
		return nil, err
	}
	cookies := NewCookies(qd.ywkey, qd.ywguid, qDInfo)
	var postData io.Reader
	var urlencoded bool
	if data != nil {
		switch data.(type) {
		case string:
			uri = uri + "?" + data.(string)
			urlencoded = true
		default:
			marshal, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}
			postData = bytes.NewBuffer(marshal)
		}
	}
	req, err := http.NewRequest(method, uri, postData)
	if data != nil {
		if urlencoded {
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		} else {
			req.Header.Set("Content-Type", "application/json")
		}
	}
	req.Header.Set("Cookie", cookies.String())
	qd.header(req)
	resp, err := qd.Cli.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(respData))
	}
	return respData, nil
}
