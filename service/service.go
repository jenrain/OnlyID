package service

import (
	"OnlyID/config"
	"OnlyID/library/log"
	"OnlyID/repository"
	"go.uber.org/zap"
)

type Service struct {
	c     *config.Config
	r     *repository.Repository
	alloc *Alloc
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
	return s
}

func (s *Service) Close() {
	s.r.Close()
}
