package service

import (
	"github.com/jenrain/OnlyID/cache"
	"github.com/jenrain/OnlyID/config"
	"github.com/jenrain/OnlyID/library/log"
	"github.com/jenrain/OnlyID/library/tool"
	"github.com/jenrain/OnlyID/repository"
	"go.uber.org/zap"
)

const (
	MYSQL = iota
	SNOWFLAKE
	REDIS
)

type Service struct {
	c         *config.Config
	R         *repository.Repository
	Cache     *cache.Redis
	alloc     *Alloc
	SnowFlake *tool.Worker
}

func NewService(c *config.Config) (s *Service) {
	var err error
	s = &Service{c: c}
	switch c.Mode {
	case MYSQL:
		s.R = repository.NewRepository(c)
		if s.alloc, err = s.NewAllocId(); err != nil {
			log.GetLogger().Error("[NewService] NewAllocId ", zap.Error(err))
			panic(err)
		}
	case SNOWFLAKE:
		if s.SnowFlake, err = s.NewAllocSnowFlakeId(); err != nil {
			log.GetLogger().Error("[NewService] NewAllocSnowFlakeId ", zap.Error(err))
			panic(err)
		}
	case REDIS:
		s.Cache = cache.NewRedisCli(c.Redis)
		if err = s.NewAllocRedisId(); err != nil {
			log.GetLogger().Error("[NewService] NewAllocRedisId ", zap.Error(err))
			panic(err)
		}
	}
	return s
}

func (s *Service) Close() {
	s.R.Close()
}
