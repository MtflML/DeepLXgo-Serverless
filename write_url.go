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
		"http://d.lfh.icu/translate",
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
