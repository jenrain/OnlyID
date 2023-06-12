package mysql

import (
	"OnlyID/library/log"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	baseLog "log"
	"os"
	"time"
)

type Config struct {
	Addr     string
	User     string
	Password string
	DbName   string

	MaxConn      int
	IdleConn     int
	QueryTimeout int //查询时间
	ExecTimeout  int //执行时间
}

func InitDB(c *Config) (db *gorm.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&readTimeout=%ds&writeTimeout=%ds",
		c.User, c.Password, c.Addr, c.DbName, c.QueryTimeout, c.ExecTimeout)

	newLogger := logger.New(
		baseLog.New(os.Stdout, "\r\n", baseLog.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.GetLogger().Error("[NewMysql] Open", zap.Any("conf", c), zap.Error(err))
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(c.MaxConn)
	sqlDB.SetMaxIdleConns(c.IdleConn)
	return
}
