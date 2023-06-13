package grpc

import (
	onlyIdSrv "OnlyID/api"
	"OnlyID/config"
	"OnlyID/library/log"
	"OnlyID/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	srv *service.Service
}

func Init(c *config.Config, s *service.Service) {
	listen, err := net.Listen("tcp", c.Server.Addr)
	if err != nil {
		panic(err)
	}
	g := grpc.NewServer()
	onlyIdSrv.RegisterOnlyIdServer(g, &Server{s})
	log.GetLogger().Info("OnlyId grpc server start", zap.Any("addr", c.Server.Addr))
	go g.Serve(listen)
}
