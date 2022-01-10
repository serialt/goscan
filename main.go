package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "gitea.com/serialt/goscan/util/conf"
	_ "gitea.com/serialt/goscan/util/logger"

	Gnsq "gitea.com/serialt/goscan/util/nsq"

	"gitea.com/serialt/goscan/util/nsq/portscan"
	"gitea.com/serialt/goscan/util/nsq/subdomain"
	"gitea.com/serialt/goscan/util/nsq/webInfo"

	"github.com/gogf/gf/frame/g"
	_ "github.com/projectdiscovery/fdmax/autofdmax"
)

func main() {
	if g.Cfg().GetBool("Domain.Enabled") { // 子域名扫描
		log.Println("[+] 子域名模块开启")
		subdomain.InitConsumer(Gnsq.SubDomainTopic, Gnsq.SubDomainChanl)
	}

	if g.Cfg().GetBool("PortScan.Enabled") { // 端口扫描
		log.Println("[+] 端口扫描模块开启")
		portscan.Init()
	}

	if g.Cfg().GetBool("WebInfo.Enabled") { // web探测
		log.Println("[+] web探测模块开启")
		webInfo.InitConsumer(Gnsq.WebInfoTopic, Gnsq.WebInfoChanl)
	}

	c := make(chan os.Signal)        // 定义一个信号的通道
	signal.Notify(c, syscall.SIGINT) // 转发键盘中断信号到c
	<-c                              // 阻塞
}
