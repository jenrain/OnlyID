package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jenrain/OnlyID/config"
	"github.com/jenrain/OnlyID/service"
)

var s *service.Service

func InitRouter(r *gin.Engine) {
	s = service.NewService(config.Conf)
	r.GET("/getId", GetId)
	r.GET("/getSnowFlakeId", GetSnowFlakeId)
	r.GET("/getRedisId", GetRedisId)
}
