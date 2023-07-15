package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jenrain/OnlyID/config"
	"github.com/jenrain/OnlyID/library/http"
	"github.com/jenrain/OnlyID/library/ip"
	"github.com/jenrain/OnlyID/library/log"
	"github.com/jenrain/OnlyID/library/tool"
	"github.com/jenrain/OnlyID/server/grpc"
	"github.com/jenrain/OnlyID/service"
)

// local ip add
var ipAddr string

func init() {
	ipAddr = ip.GetIpAddr()
	fmt.Println("local ip addr: ", ipAddr)
}

func main() {
	flag.Parse()
	if err := config.Init(); err != nil {
		panic(err)
	}
	log.NewLogger(config.Conf.Log)
	s := service.NewService(config.Conf)
	grpc.Init(config.Conf, s)
	if err := tool.InitMasterNode(config.Conf.Etcd, ipAddr+":"+config.Conf.Server.Port, 5); err != nil {
		panic(err)
	}
	// http服务 测试用
	r := gin.Default()
	http.InitRouter(r)
	r.Run(":8083")
	// 优雅关闭
	tool.QuitSignal(func() {
		s.Close()
		log.GetLogger().Info("OnlyId exit success")
	})
}
