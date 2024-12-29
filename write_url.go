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
		"https://deeplx.smnet.io/translate",
		"http://d.lfh.icu/translate",
		"https://deeplx-2dn.pages.dev:443/translate",
		"https://deeplx.mingming.dev:443/translate",
		"http://142.171.225.251:1188/translate",
		"http://107.172.87.53:1188/translate",
		"http://45.91.81.18:1188/translate",
		"http://150.230.46.219:8136/translate",
		"http://158.101.137.14:1188/translate",
		"http://51.81.187.241:1188/translate",
		"http://211.227.72.101:1188/translate",
		"http://158.101.148.105:1188/translate",
		"http://138.2.11.53:1188/translate",
		"http://132.226.232.50:1188/translate",
		"http://123.56.13.17:1188/translate",
		"https://deeplx-serverless.vercel.app/translate",
		"http://106.14.155.58:18019/translate",
		"http://47.76.34.130:20000/translate",
		"http://156.254.115.38:20000/translate",
		"http://107.172.87.53:1188/translate",
		"http://20.188.101.121:9000/translate",
		"https://d.lfh.icu/translate",
		"https://deeplx.edu6.eu.org/translate?token=666666",
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
