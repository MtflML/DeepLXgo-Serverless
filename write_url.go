package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// GenerateURLList 创建一个包含多个URL的列表
func GenerateURLList() []string {
	urlList := []string{
		"https://api.deeplx.org/_pZpvSNySwhVy1DYkd6tSnauK8Dxj1IOPgAvP-TGgNE/translate",
		"http://38.175.202.113:1188/translate",
		"https://mikeee-deeplx.hf.space/translate",
		"https://deeplx.dattw.eu.org/translate",
		"http://142.171.19.149:9000/translate",
		"http://138.68.240.43:1188/translate",
		"https://deeplx.smnet.io/translate",
		"http://142.171.18.103:1188/translate",
		"https://deeplx.smnet.io/translate",
		"https://deeplx.doi9.top/translate",
		"https://deeplx.qninq.cn/translate",
		"https://deeplx.niubipro.com/translate",
		"https://deeplx-serverless.vercel.app/translate",
		"http://4.218.9.222:9000/translate",
		"http://150.109.240.249:9000/translate",
		"http://172.83.158.171:1188/translate",
		"http://45.145.72.182:1188/translate",
		"http://172.83.157.240:1188/translate",
		"http://45.145.75.161:1188/translate",
		"http://172.83.158.132:1188/translate",
		"http://45.145.72.146:1188/translate",
		"http://45.145.74.178:1188/translate",
		"http://172.83.158.181:1188/translate",
		"http://47.120.44.206:9000/translate",
		"http://45.143.233.5:9000/translate",
		"https://deeplx.edu6.eu.org/translate?token=666666",
		"https://deeplx.030101.xyz/api",
		"https://deeplx.missuo.ru/translate?key=-RRuIvlL7QBF5TrwSc62fQ_LnPkCmxfSstxaAM69iZI=",
	}
	
	return urlList
}

// WriteURLListToFile 将URL列表写入到指定文件
func WriteURLListToFile(urlList []string, filePath string) error {
	var data strings.Builder
	for _, url := range urlList {
		data.WriteString(url + "\n")
	}

	if err := ioutil.WriteFile(filePath, []byte(data.String()), 0644); err != nil {
		return fmt.Errorf("failed to write URL list to file: %w", err)
	}
	fmt.Printf("URL list written to file '%s'\n", filePath)
	return nil
}
