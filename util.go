package qdapi

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

func GetProxyClient() *http.Client {
	address := "localhost:8888"
	conn, err := net.DialTimeout("tcp", address, time.Second*2)
	if err != nil {
		return http.DefaultClient
	}
	defer conn.Close()
	//for Charles
	proxyURL, err := url.Parse("http://" + address)
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
