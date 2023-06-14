package cache

import (
	"context"
	"github.com/jenrain/OnlyID/library/log"
	"go.uber.org/zap"
)

type Redis struct {
	opt      *RedisClientOption
	redisCli *RedisClient
}

func NewRedisCli(opt *RedisClientOption) (cli *Redis) {
	cli = &Redis{
		opt:      opt,
		redisCli: NewRedisClient(opt),
	}
	return
}

func (r *Redis) Ping() (string, error) {
	return r.redisCli.Ping(context.TODO())
}

func (r *Redis) GetId(bizTag string) (int64, error) {
	res, err := r.redisCli.Incr(context.TODO(), bizTag)
	if err != nil {
		log.GetLogger().Error("[GetId] Incr", zap.Any("bizTag", bizTag), zap.Error(err))
		return -1, err
	}
	return int64(res), nil
}

func (r *Redis) Close() {
	r.redisCli.connPool.Close()
	log.GetLogger().Info("redis pool successfully closed")
}
