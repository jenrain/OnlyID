package service

import "errors"

func (s *Service) NewAllocRedisId() error {
	ping, _ := s.Cache.Ping()
	if ping != "PONG" {
		return errors.New("connect to redis  fail")
	}
	return nil
}

func (s *Service) RedisGetId(bizTag string) (int64, error) {
	return s.Cache.GetId(bizTag)
}
