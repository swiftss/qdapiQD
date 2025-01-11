package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
)

func GetProxyClient() *http.Client {
	//for Charles
	proxyURL, err := url.Parse("http://localhost:8888")
	if err != nil {
		log.Fatal("Invalid proxy URL:", err)
	}
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}
}
func GetInsecureClient() *http.Client {
	//for Charles
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // 忽略证书验证
		},
	}
	return &http.Client{
		Transport: tr,
	}
}
