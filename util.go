package qdapi

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
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

func LoadConfigFromJSON(filename string) ([]QiDianApiConfig, error) {
	var config []QiDianApiConfig
	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return config, err
}

func SaveConfigToJSON(filename string, data interface{}) error {
	// 创建文件
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 创建JSON编码器
	encoder := json.NewEncoder(file)

	// 设置格式缩进（可选）
	encoder.SetIndent("", "  ")

	// 执行编码
	return encoder.Encode(data)
}
