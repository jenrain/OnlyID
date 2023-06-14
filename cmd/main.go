package main

import (
	"flag"
	"github.com/jenrain/OnlyID/config"
	"github.com/jenrain/OnlyID/library/log"
	"github.com/jenrain/OnlyID/library/tool"
	"github.com/jenrain/OnlyID/server/grpc"
	"github.com/jenrain/OnlyID/service"
)

func main() {
	flag.Parse()
	if err := config.Init(); err != nil {
		panic(err)
	}
	log.NewLogger(config.Conf.Log)
	s := service.NewService(config.Conf)
	grpc.Init(config.Conf, s)
	if err := tool.InitMasterNode(config.Conf.Etcd, config.Conf.Server.Addr, 30); err != nil {
		panic(err)
	}
	// 优雅关闭
	tool.QuitSignal(func() {
		s.Close()
		log.GetLogger().Info("OnlyId exit success")
	})
}
