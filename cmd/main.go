package main

import (
	"OnlyID/config"
	"OnlyID/library/log"
	"OnlyID/library/tool"
	"OnlyID/server/grpc"
	"OnlyID/service"
	"flag"
)

func main() {
	flag.Parse()
	if err := config.Init(); err != nil {
		panic(err)
	}
	log.NewLogger(config.Conf.Log)
	s := service.NewService(config.Conf)
	grpc.Init(config.Conf, s)

	// 优雅关闭
	tool.QuitSignal(func() {
		s.Close()
		log.GetLogger().Info("OnlyId exit success")
	})
}
