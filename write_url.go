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
		"https://deeplx-2dn.pages.dev:443/translate",
		"https://deeplx.mingming.dev:443/translate",
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
