package main

import (
	"context"
	"deeplx-local/domain"
	"fmt"
	"github.com/imroc/req/v3"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"golang.org/x/sync/errgroup"
	"syscall"
	"time"
)

var (
	client   = req.NewClient().SetTimeout(time.Second)
	validReq = domain.TranslateRequest{
		Text:       "I love you",
		SourceLang: "EN",
		TargetLang: "ZH",
	}
)

// getValidURLs 从文件中读取并处理URL，返回可用URL列表
func getValidURLs() []string {
	content, err := os.ReadFile("url.txt")
	if err != nil {
		log.Fatal(err)
	}

	urls := strings.Split(string(content), "\n")

	var wg sync.WaitGroup
	var mutex sync.Mutex
	validURLs := make([]string, 0)

	checkURL := func(url string) {
		defer wg.Done()

		if !strings.HasSuffix(url, "/translate") {
			url += "/translate"
		}
		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}

		availability, err := checkURLAvailability(url)
		if err != nil {
			log.Printf("Error checking URL %s: %s\n", url, err)
			return
		}

		if availability {
			mutex.Lock()
			validURLs = append(validURLs, url)
			mutex.Unlock()
		}
	}

	wg.Add(len(urls))
	for _, url := range urls {
		go checkURL(url)
	}

	wg.Wait()

	log.Printf("Available URLs count: %d\n", len(validURLs))
	return validURLs
}

// checkURLAvailability 检查URL是否可用
func checkURLAvailability(url string) (bool, error) {
	var result domain.TranslateResponse
	response, err := client.R().SetBody(&validReq).SetSuccessResult(&result).Post(url)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	return "我爱你" == result.Data, nil
}

// exit 监听退出信号并优雅关闭HTTP服务器
func exit(engine *http.Server) {
	osSig := make(chan os.Signal, 1)
	quit := make(chan bool, 1)

	signal.Notify(osSig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		sig := <-osSig
		fmt.Println("Received exit signal:", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := engine.Shutdown(ctx); err != nil {
			fmt.Println("Web service shutdown failed:", err)
		}
		quit <- true
	}()

	<-quit
	fmt.Println("Service PID:", os.Getpid())
	fmt.Println("Service has exited.")

	// 查杀
	exec.Command("killall", "main", strconv.Itoa(os.Getpid())).Run()
	// 自杀
	exec.Command("kill", "-9", strconv.Itoa(os.Getpid())).Run()
}
