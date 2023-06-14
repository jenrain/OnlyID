package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/jenrain/OnlyID/cache"
	"github.com/jenrain/OnlyID/library/database/mysql"
	"github.com/jenrain/OnlyID/library/log"
)

var (
	confPath string
	Conf     = new(Config)
)

type Config struct {
	Development bool
	SnowFlakeId int64
	Etcd        []string
	Log         *log.Options
	Mysql       *mysql.Config
	Redis       *cache.RedisClientOption
	Server      *Srv
	Mode        int
}

type Srv struct {
	Addr string
}

func init() {
	flag.StringVar(&confPath, "conf", "./only_id.toml", "default config path")
}

func Init() error {
	return local()
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}
