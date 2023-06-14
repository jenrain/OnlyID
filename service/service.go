package service

import (
	"github.com/jenrain/OnlyID/config"
	"github.com/jenrain/OnlyID/library/log"
	"github.com/jenrain/OnlyID/library/tool"
	"github.com/jenrain/OnlyID/repository"
	"go.uber.org/zap"
)

type Service struct {
	c         *config.Config
	r         *repository.Repository
	alloc     *Alloc
	snowFlake *tool.Worker
}

func NewService(c *config.Config) (s *Service) {
	var err error
	s = &Service{
		c: c,
		r: repository.NewRepository(c),
	}
	if s.alloc, err = s.NewAllocId(); err != nil {
		log.GetLogger().Error("[NewService] NewAllocId ", zap.Error(err))
		panic(err)
	}
	if s.snowFlake, err = s.NewAllocSnowFlakeId(); err != nil {
		log.GetLogger().Error("[NewService] NewAllocSnowFlakeId ", zap.Error(err))
		panic(err)
	}
	return s
}

func (s *Service) Close() {
	s.r.Close()
}
