package main

import (
	"deeplx-local/service"
	"deeplx-local/web"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	“os"
)

func main() {
	if err := WriteURLListToFile(GenerateURLList(), "url.txt"); err != nil {
		log.Fatalf("Error writing URL list: %v", err)
	}

	log.Println("url.txt generated successfully.")

	// 从文件中读取并处理URL
	urls := getValidURLs()

	// 注册服务
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	lxService := service.NewDeepLXService(&urls)
	balancerService := service.NewLoadBalancer(lxService.(*service.DeepLXService))
	lxHandler := web.NewDeepLXHandler(balancerService)
	lxHandler.RegisterRoutes(r)

	// 启动服务
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	go func() {
		if err := http.ListenAndServe(":"+port, r); err != nil {
			log.Fatal("Failed to start web service:", err)
		}
	}()
	// 监听退出
	exit()
}
