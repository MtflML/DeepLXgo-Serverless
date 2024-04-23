package main

import (
	"deeplx-local/service"
	"deeplx-local/web"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("url.txt")
	if err != nil {
		fmt.Println("创建文件时发生错误:", err)
		return
	}
	defer file.Close()

	lines := []string{
		"https://api.deeplx.org/translate",
		"https://deeplxpro.vercel.app/translate",
		"https://deeplx.llleman.com/translate",
		"https://deepl.tr1ck.cn/translate",
		"https://deeplx.ychinfo.com/translate",
		"https://dx-api.nosec.link/translate",
		"https://free-deepl.speedcow.top/translate",
		}
	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			fmt.Println("写入文件时发生错误:", err)
			return
		}
	}

	fmt.Println("写入多行成功")
	
	// 从文件中读取并处理URL
	urLs := getValidURLs()

	// 注册服务
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	lxService := service.NewDeepLXService(&urLs)
	balancerService := service.NewLoadBalancer(lxService.(*service.DeepLXService))
	lxHandler := web.NewDeepLXHandler(balancerService)
	lxHandler.RegisterRoutes(r)

	// 启动服务
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := &http.Server{
		Addr:    ":"+port,
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println("web服务启动失败: ", err)
		}
	}()

	// 监听退出
	exit(server)
}
