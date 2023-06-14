package cache

import (
	"context"
	"errors"
	"github.com/gomodule/redigo/redis"
	"github.com/jenrain/OnlyID/library/log"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type RedisClientOption struct {
	Addr            string
	Password        string
	ConnTimeout     int
	ReadTimeout     int
	WriteTimeout    int
	MaxIdle         int
	IdleTimeout     int
	MaxActive       int
	MaxConnLifetime int
}

type RedisClient struct {
	Opt      *RedisClientOption
	connPool *redis.Pool
}

func newPool(opt *RedisClientOption) *redis.Pool {
	pool := &redis.Pool{
		IdleTimeout: time.Duration(opt.IdleTimeout) * time.Millisecond,
		// Maximum number of idle connections in the pool.
		// configuring a connection.
		MaxConnLifetime: time.Duration(opt.MaxConnLifetime) * time.Millisecond,
		MaxIdle:         opt.MaxIdle,
		MaxActive:       opt.MaxActive,
		// Dial is an application supplied function for creating and
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",
				opt.Addr,
				redis.DialConnectTimeout(time.Duration(opt.ConnTimeout)*time.Millisecond),
				redis.DialReadTimeout(time.Duration(opt.ReadTimeout)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(opt.WriteTimeout)*time.Millisecond),
				redis.DialPassword(opt.Password),
			)
		},
	}
	return pool
}

func NewRedisClient(opt *RedisClientOption) *RedisClient {
	pool := newPool(opt)

	if pool == nil {
		log.GetLogger().Error("[NewRedisClient] Init Redis Client")
	}

	return &RedisClient{
		connPool: pool,
	}
}

func (rc *RedisClient) GetActive() int {
	return rc.connPool.ActiveCount()
}

func (rc *RedisClient) GetIdleCount() int {
	return rc.connPool.IdleCount()
}

func (rc *RedisClient) Ping(ctx context.Context) (string, error) {
	conn := rc.connPool.Get()
	defer conn.Close()
	return redis.String(conn.Do("Ping"))
}

func (rc *RedisClient) Incr(ctx context.Context, key string) (int, error) {
	conn := rc.connPool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("INCR", key))
}

func (rc *RedisClient) Expire(ctx context.Context, key string, ttlSec int) (int, error) {
	conn := rc.connPool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("EXPIRE", key, ttlSec))
}

func (rc *RedisClient) Exists(ctx context.Context, key string) (ex bool, err error) {
	var reply int64
	conn := rc.connPool.Get()
	defer conn.Close()
	reply, err = redis.Int64(conn.Do("Exists", key))
	if reply == 1 {
		ex = true
		return
	}
	return
}

func (rc *RedisClient) Set(ctx context.Context, args ...interface{}) (string, error) {
	conn := rc.connPool.Get()
	defer conn.Close()
	return redis.String(conn.Do("SET", args...))
}

func (rc *RedisClient) SetNxEx(ctx context.Context, key string, value string, timeout int64) (string, error) {
	conn := rc.connPool.Get()
	defer conn.Close()
	return redis.String(conn.Do("SET", key, value, "NX", "EX", timeout))
}

func (rc *RedisClient) SAdd(ctx context.Context, key string, values ...interface{}) (int, error) {
	conn := rc.connPool.Get()
	defer conn.Close()
	if len(key) == 0 {
		log.GetLogger().Error("SAdd", zap.Any("args", values))
	}
	args := append([]interface{}{key}, values...)
	return redis.Int(conn.Do("SADD", args...))
}

func (rc *RedisClient) SMembers(ctx context.Context, key string) ([]string, error) {
	conn := rc.connPool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("SMEMBERS", key))
}

func (rc *RedisClient) SPop(ctx context.Context, key string) (string, error) {
	conn := rc.connPool.Get()
	defer conn.Close()
	return redis.String(conn.Do("SPOP", key))
}

func (rc *RedisClient) MGet(ctx context.Context, keys ...interface{}) ([]string, error) {
	conn := rc.connPool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("MGET", keys...))
}

func (rc *RedisClient) SRem(ctx context.Context, key string, members ...interface{}) (int, error) {
	conn := rc.connPool.Get()
	defer conn.Close()
	args := append([]interface{}{key}, members...)
	return redis.Int(conn.Do("SREM", args...))
}

func (rc *RedisClient) SCard(ctx context.Context, key string) (int, error) {
	conn := rc.connPool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("SCARD", key))
}

func (rc *RedisClient) Append(ctx context.Context, key string, val string) (int, error) {
	conn := rc.connPool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("APPEND", key, val))
}

func (rc *RedisClient) IncrByFloat(ctx context.Context, key string, val float64) (string, error) {
	conn := rc.connPool.Get()
	defer conn.Close()
	return redis.String(conn.Do("INCRBYFLOAT", key, val))
}

func (rc *RedisClient) IncrBy(ctx context.Context, key string, val int) (int, error) {
	conn := rc.connPool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("INCRBY", key, val))
}

func (rc *RedisClient) Get(ctx context.Context, key string) (string, error) {
	conn := rc.connPool.Get()
	defer conn.Close()
	return redis.String(conn.Do("GET", key))
}

func (rc *RedisClient) Setnx(ctx context.Context, key, val string) (int, error) {
	conn := rc.connPool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("SETNX", key, val))
}

func (rc *RedisClient) Del(ctx context.Context, key string) (int, error) {
	conn := rc.connPool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("Del", key))
}

func (rc *RedisClient) ZRange(ctx context.Context, key string, start, stop int) ([]string, error) {
	conn := rc.connPool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("ZRANGE", key, start, stop))
}

func (rc *RedisClient) ZScore(ctx context.Context, key, val string) (int64, error) {
	conn := rc.connPool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("ZSCORE", key, val))
}

func (rc *RedisClient) ZRangeWithScores(ctx context.Context, key string, start, stop int) (map[string]string, error) {
	conn := rc.connPool.Get()
	defer conn.Close()

	reply, err := redis.StringMap(conn.Do("ZRANGE", key, start, stop, "WITHSCORES"))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (rc *RedisClient) ZRangeByScore(ctx context.Context, key string, min, max int64, minopen, maxopen bool) ([]string, error) {
	c := rc.connPool.Get()
	defer c.Close()

	minstr := strconv.FormatInt(int64(min), 10)
	maxstr := strconv.FormatInt(int64(max), 10)
	if minopen {
		minstr = "(" + strconv.FormatInt(int64(min), 10)
	}

	if maxopen {
		maxstr = "(" + strconv.FormatInt(int64(max), 10)
	}

	if 0 == min {
		minstr = "-inf"
	}

	if -1 == max {
		maxstr = "+inf"
	}

	reply, err := redis.Strings(c.Do("ZRANGEBYSCORE", key, minstr, maxstr))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (rc *RedisClient) ZRangeByScoreWithScores(ctx context.Context, key string, min, max int, minopen, maxopen bool) (map[string]string, error) {
	c := rc.connPool.Get()
	defer c.Close()

	minstr := strconv.FormatInt(int64(min), 10)
	maxstr := strconv.FormatInt(int64(max), 10)
	if minopen {
		minstr = "(" + strconv.FormatInt(int64(min), 10)
	}

	if maxopen {
		maxstr = "(" + strconv.FormatInt(int64(max), 10)
	}

	if 0 == min {
		minstr = "-inf"
	}

	if -1 == max {
		maxstr = "+inf"
	}

	reply, err := redis.StringMap(c.Do("ZRANGEBYSCORE", key, minstr, maxstr, "WITHSCORES"))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (rc *RedisClient) HGetall(ctx context.Context, key string) (map[string]string, error) {
	c := rc.connPool.Get()
	defer c.Close()

	reply, err := redis.StringMap(c.Do("HGETALL", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (rc *RedisClient) HINCRBY(ctx context.Context, key, subkey string, val int64) (err error) {
	c := rc.connPool.Get()
	defer c.Close()

	_, err = redis.Int64(c.Do("HINCRBY", key, subkey, val))
	return
}

func (rc *RedisClient) HGet(ctx context.Context, key, subkey string) (string, error) {
	c := rc.connPool.Get()
	defer c.Close()

	reply, err := redis.String(c.Do("HGET", key, subkey))
	if err != nil {
		return "", err
	}

	return reply, nil
}

func (rc *RedisClient) HSet(ctx context.Context, key, subkey, value string) error {
	c := rc.connPool.Get()
	defer c.Close()

	_, err := redis.Int(c.Do("HSET", key, subkey, value))
	if err != nil {
		return err
	}

	return nil
}

func (rc *RedisClient) SetEx(ctx context.Context, key, value string, ttl int64) error {
	c := rc.connPool.Get()
	defer c.Close()

	reply, err := redis.String(c.Do("SETEX", key, ttl, value))
	if err != nil {
		return err
	}

	if reply == "OK" {
		return nil
	} else {
		return errors.New("redisclient: unexpected reply of setex")
	}
}

func (rc *RedisClient) ZRemRangeByScore(ctx context.Context, key string, min, max int64) (reply int64, err error) {
	c := rc.connPool.Get()
	defer c.Close()

	minstr := strconv.FormatInt(min, 10)
	maxstr := strconv.FormatInt(max, 10)
	if 0 == min {
		minstr = "-inf"
	}
	if -1 == max {
		maxstr = "+inf"
	}
	return redis.Int64(c.Do("ZREMRANGEBYSCORE", key, minstr, maxstr))
}

func (rc *RedisClient) ZRemRangeByRank(ctx context.Context, key string, min, max int64) (reply int64, err error) {
	c := rc.connPool.Get()
	defer c.Close()
	minstr := strconv.FormatInt(min, 10)
	maxstr := strconv.FormatInt(max, 10)
	return redis.Int64(c.Do("ZREMRANGEBYRANK", key, minstr, maxstr))
}

func (rc *RedisClient) ZCount(ctx context.Context, key string, min, max int64) (reply int64, err error) {
	c := rc.connPool.Get()
	defer c.Close()
	minstr := strconv.FormatInt(min, 10)
	maxstr := strconv.FormatInt(max, 10)
	if 0 == min {
		minstr = "-inf"
	}
	if -1 == max {
		maxstr = "+inf"
	}
	return redis.Int64(c.Do("ZCOUNT", key, minstr, maxstr))
}

func (rc *RedisClient) ZRem(ctx context.Context, keys ...interface{}) (reply int64, err error) {
	c := rc.connPool.Get()
	defer c.Close()
	return redis.Int64(c.Do("ZREM", keys...))
}

func (rc *RedisClient) HKeys(ctx context.Context, key string) (replys []string, err error) {
	c := rc.connPool.Get()
	defer c.Close()
	return redis.Strings(c.Do("HKEYS", key))
}

func (rc *RedisClient) HMSet(ctx context.Context, key string, values ...interface{}) (err error) {
	c := rc.connPool.Get()
	defer c.Close()

	params := make([]interface{}, 0)
	params = append(params, key)
	params = append(params, values...)
	_, err = redis.String(c.Do("HMSET", params...))
	return
}

func (rc *RedisClient) HMGet(ctx context.Context, key string, fields []interface{}) (replys []string, err error) {
	c := rc.connPool.Get()
	defer c.Close()
	params := make([]interface{}, 0)
	params = append(params, key)
	params = append(params, fields...)

	return redis.Strings(c.Do("HMGET", params...))
}

func (rc *RedisClient) HDel(ctx context.Context, key string, fields ...interface{}) (reply int64, err error) {
	c := rc.connPool.Get()
	defer c.Close()
	params := make([]interface{}, 0)
	params = append(params, key)
	params = append(params, fields...)
	reply, err = redis.Int64(c.Do("HDEL", params...))
	return
}

func (rc *RedisClient) SisMember(ctx context.Context, key string, subKey string) (result bool, err error) {
	c := rc.connPool.Get()
	defer c.Close()
	result, err = redis.Bool(c.Do("SISMEMBER", key, subKey))
	return
}

// ZAdd key score member [[score member] [score member] ...]
func (rc *RedisClient) ZAdd(ctx context.Context, key string, args ...interface{}) (err error) {
	c := rc.connPool.Get()
	defer c.Close()
	_, err = redis.Int(c.Do("ZADD", append([]interface{}{interface{}(key)}, args...)...))
	return
}

func (rc *RedisClient) RunScript(ctx context.Context, keyCount int, script string, args ...interface{}) (n int, err error) {
	c := rc.connPool.Get()
	defer c.Close()

	luaScript := redis.NewScript(keyCount, script)
	return redis.Int(luaScript.Do(c, args...))
}

// ------------------辅助函数--------------------//
const (
	CodisMetricsUseCaseName = "use_case"
	UseCaseFeatureProduct   = "feature_product"
	UseCaseWriteHistory     = "write_history"
	UseCaseMergeKey         = "merge_key"
	UseCaseCoordFreqReduce  = "coord_freq_reduce"
	UseCaseCycleKey         = "cycle_key"
	UseCaseEventPool        = "event_pool"
)

// redis模块 context相关函数
func NewContextWithValue(ctx context.Context, key string, value string) context.Context {
	if ctx == nil {
		return nil
	}
	return context.WithValue(ctx, key, value)
}

func ContextValue(ctx context.Context, key string) string {
	if ctx == nil {
		return ""
	}
	val, _ := ctx.Value(key).(string)
	return val
}
