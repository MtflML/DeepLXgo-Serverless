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
		"http://45.125.18.25:1188/translate",
		"http://8.130.41.212:1188/translate",
		"http://120.55.63.151:1188/translate",
		"http://101.43.224.133:1188/translate",
		"http://119.91.23.165:1188/translate",
		"http://121.43.36.212:1188/translate",
		"http://165.154.5.156:1188/translate",
		"http://109.206.245.73:1188/translate",
		"http://121.41.99.203:1188/translate",
		"http://8.130.123.166:1188/translate",
		"http://43.132.160.253:1188/translate",
		"http://159.75.240.245:1188/translate",
		"http://43.139.108.188:1188/translate",
		"http://47.76.126.155:1188/translate",
		"http://42.194.148.47:1188/translate",
		"http://59.110.162.193:1188/translate",
		"http://101.42.160.235:1188/translate",
		"http://81.70.207.129:1188/translate",
		"http://47.119.24.34:1188/translate",
		"http://116.252.28.28:1188/translate",
		"http://8.130.135.21:1188/translate",
		"http://194.104.146.32:1188/translate",
		"http://8.210.238.191:1188/translate",
		"http://47.242.206.190:1188/translate",
		"http://139.224.191.20:1188/translate",
		"http://112.126.74.237:1188/translate",
		"http://47.243.57.229:1188/translate",
		"http://8.138.111.57:1188/translate",
		"http://116.62.112.61:1188/translate",
		"http://47.76.92.38:1188/translate",
		"http://123.116.149.203:1188/translate",
		"http://47.119.172.46:1188/translate",
		"http://120.78.15.169:1188/translate",
		"http://124.223.85.170:1188/translate",
		"http://120.26.116.45:1188/translate",
		"http://139.159.242.136:1188/translate",
		"http://114.132.74.10:1188/translate",
		"http://8.131.60.61:1188/translate",
		"http://47.102.106.172:1188/translate",
		"http://113.89.232.208:1188/translate",
		"http://58.219.74.154:1188/translate",
		"http://182.148.64.26:1188/translate",
		"http://43.154.208.219:1188/translate",
		"http://43.143.233.18:1188/translate",
		"http://39.98.115.99:1188/translate",
		"http://47.109.98.57:1188/translate",
		"http://120.55.189.173:1188/translate",
		"http://82.157.49.14:1188/translate",
		"http://119.28.13.127:1188/translate",
		"http://119.91.153.247:1188/translate",
		"http://124.222.168.113:1188/translate",
		"http://122.152.221.70:1188/translate",
		"http://116.23.76.92:1188/translate",
		"http://103.247.28.143:1188/translate",
		"http://42.192.83.104:1188/translate",
		"http://175.178.54.19:1188/translate",
		"http://42.193.219.103:1188/translate",
		"http://39.105.97.161:1188/translate",
		"http://60.205.231.83:1188/translate",
		"http://8.142.79.40:1188/translate",
		"http://101.43.16.48:1188/translate",
		"http://8.134.151.104:1188/translate",
		"http://150.158.45.100:1188/translate",
		"http://39.101.74.119:1188/translate",
		"http://101.133.228.48:1188/translate",
		"http://42.192.21.116:1188/translate",
		"http://106.14.72.237:1188/translate",
		"http://124.71.99.68:1188/translate",
		"http://121.43.134.47:1188/translate",
		"http://103.152.35.2:1188/translate",
		"http://101.43.76.234:1188/translate",
		"http://120.76.141.173:1188/translate",
		"http://47.107.100.134:1188/translate",
		"http://116.48.95.229:1188/translate",
		"http://47.76.48.116:1188/translate",
		"http://43.129.77.32:1188/translate",
		"http://123.57.17.229:1188/translate",
		"http://39.105.60.208:1188/translate",
		"http://39.107.101.134:1188/translate",
		"http://1.12.45.103:1188/translate",
		"http://139.224.225.116:1188/translate",
		"http://47.103.194.159:1188/translate",
		"http://39.100.95.114:1188/translate",
		"http://45.152.67.153:1188/translate",
		"http://190.92.242.194:1188/translate",
		"http://47.122.75.223:1188/translate",
		"http://124.220.101.194:1188/translate",
		"http://35.220.175.223:1188/translate",
		"http://101.132.151.32:1188/translate",
		"http://43.143.240.2:1188/translate",
		"http://120.78.82.181:1188/translate",
		"http://101.133.234.93:1188/translate",
		"http://81.70.94.105:1188/translate",
		"http://124.71.191.52:1188/translate",
		"http://120.76.140.44:1188/translate",
		"http://123.56.13.17:1188/translate",
		"http://139.159.254.35:1188/translate",
		"http://124.221.198.187:1188/translate",
		"http://42.98.172.229:1188/translate",
		"http://8.130.121.171:1188/translate",
		"http://101.201.38.103:1188/translate",
		"http://120.46.95.125:1188/translate",
		"http://47.92.227.110:1188/translate",
		"http://43.139.218.213:1188/translate",
		"http://1.14.59.89:1188/translate",
		"http://106.14.104.93:1188/translate",
		"http://47.107.109.131:1188/translate",
		"http://120.79.93.103:1188/translate",
		"http://119.28.32.110:1188/translate",
		"http://43.143.198.166:1188/translate",
		"http://8.140.246.236:1188/translate",
		"http://116.204.90.243:2001/translate",
		"http://47.238.223.93/translate",
		"https://49.232.185.166/translate",
		"http://150.158.80.45:6009/translate",
		"http://116.231.112.27:1188/translate",
		"http://222.69.212.47:1188/translate",
		"http://59.110.34.163:85/translate",
		"https://dx.mmyy.fun/translate",
		"http://dx.mmyy.fun/translate",
		"https://154.18.161.26/translate",
		"https://59.174.227.48:1024/translate",
		"http://154.18.161.26/translate",
		"http://deepl.kevinzhang.cn/translate",
		"https://deepl.kevinzhang.cn/translate",
		"http://16.163.29.52:12294/translate",
		"http://101.35.129.85:29999/translate",
		"http://deepl.tr1ck.cn/translate",
		"https://deepl.tr1ck.cn/translate",
		"http://106.15.89.0:7001/translate",
		"http://153.36.242.81:10030/translate",
		"https://deeplx.iloli.love/translate",
		"http://deeplx.iloli.love/translate",
		"http://dx.ift.lat/translate",
		"https://dx.ift.lat/translate",
		"http://38.207.175.96:8082/translate",
		"https://deepx.api.neix.in/translate",
		"http://deepx.api.neix.in/translate",
		"http://45.94.43.74:8081/translate",
		"http://45.94.43.74:8082/translate",
		"http://112.193.113.149:9527/translate",
		"http://110.42.235.245:8811/translate",
		"http://82.156.183.23:6000/translate",
		"https://8.210.6.88/translate",
		"https://translate.dftianyi.com/translate",
		"https://e.nxnow.top/translate",
		"https://bvhk.sftp.nxnow.top/translate",
		"http://1.14.71.133:1188/translate",
		"http://119.91.152.74:1188/translate",
		"http://1.12.243.147:1188/translate",
		"http://124.222.50.132:1188/translate",
		"http://103.152.34.47:1188/translate",
		"http://124.223.210.72:1188/translate",
		"http://117.50.183.46:1188/translate",
		"http://106.14.17.223:1188/translate",
		"http://123.60.157.70:8085/translate",
		"http://116.204.118.161:1188/translate",
		"http://175.178.237.179:1188/translate",
		"http://134.122.133.56:1188/translate",
		"http://180.164.162.35:1188/translate",
		"http://134.122.133.76:1188/translate",
		"http://134.122.133.65:1188/translate",
		"http://203.25.119.208:1188/translate",
		"http://152.67.211.94:1188/translate",
		"http://107.175.28.239:1188/translate",
		"http://152.70.117.91:1188/translate",
		"http://142.171.172.172:1188/translate",
		"http://107.173.148.148:1188/translate",
		"http://168.138.34.126:1188/translate",
		"http://148.135.107.108:1188/translate",
		"http://142.171.184.251:1188/translate",
		"http://104.234.60.178:1188/translate",
		"http://146.190.200.191:1188/translate",
		"http://107.174.67.205:1188/translate",
		"http://107.191.39.37:1188/translate",
		"http://35.229.248.32:1188/translate",
		"http://43.138.68.36:8085/translate",
		"http://162.55.35.20:1188/translate",
		"http://18.142.80.110:1188/translate",
		"http://43.154.115.62:1188/translate",
		"http://130.61.194.102:1188/translate",
		"http://116.203.83.80:1188/translate",
		"http://211.227.72.101:1188/translate",
		"http://37.123.196.26:1188/translate",
		"http://152.67.208.218:8000/translate",
		"http://43.133.208.75:1188/translate",
		"http://132.145.80.159:1188/translate",
		"http://168.138.214.221:1188/translate",
		"http://146.56.180.125:1188/translate",
		"http://209.141.49.210:1188/translate",
		"http://129.153.73.237:1188/translate",
		"http://132.226.14.46:1188/translate",
		"http://45.66.217.9:1188/translate",
		"http://47.243.180.227:1188/translate",
		"http://38.148.254.10:1188/translate",
		"http://146.56.97.135:1188/translate",
		"http://195.154.184.125:1188/translate",
		"http://43.153.38.254:1188/translate",
		"http://157.245.192.219:1188/translate",
		"http://165.22.247.82:1188/translate",
		"http://192.227.178.132:1188/translate",
		"http://49.232.164.78:1188/translate",
		"http://49.233.41.73:1188/translate",
		"http://49.235.73.101:1188/translate",
		"http://157.90.115.116:1188/translate",
		"http://45.145.72.29:1188/translate",
		"http://43.152.203.138:1188/translate",
		"http://207.148.127.142:1188/translate",
		"http://52.253.96.13:1188/translate",
		"http://46.148.235.162:1188/translate",
		"http://172.98.13.238:1188/translate",
		"http://64.112.42.240:1188/translate",
		"http://194.87.252.161:1188/translate",
		"http://62.210.144.227:1188/translate",
		"http://82.157.137.187:1188/translate",
		"http://82.156.36.11:1188/translate",
		"http://8.142.134.155:1188/translate",
		"http://82.157.157.107:1188/translate",
		"http://8.218.202.14:1188/translate",
		"http://91.199.209.52:1188/translate",
		"http://deeplx.hc-beijing.rhythmlian.cn/translate",
		"https://deepl.degbug.top/translate",
		"http://74.48.17.203:1188/translate",
		"http://72.18.80.51:1188/translate",
		"https://deep.jftkj.cyou/translate",
		"http://77.169.34.84:1188/translate",
		"https://deepl.zhaosaipo.com/translate",
		"https://ghhosa.zzaning.com/translate",
		"http://hc-beijing.rhythmlian.cn:8085/translate",
		"https://deeplx.keyrotate.com/translate",
		"https://deepx.dumpit.top/translate",
		"http://46.8.19.173:1188/translate",
		"http://fenghua.site:1188/translate",
		"https://api.deeplx.org/translate",
		"https://deepl.arthals.ink/translate",
		"http://deepl.arthals.ink/translate"
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
