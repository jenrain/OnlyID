package repository

import (
	"github.com/jenrain/OnlyID/config"
	"github.com/jenrain/OnlyID/library/database/mysql"
	"github.com/jenrain/OnlyID/library/log"
	"gorm.io/gorm"
)

type Repository struct {
	c  *config.Config
	db *gorm.DB
}

func NewRepository(c *config.Config) (r *Repository) {
	r = &Repository{
		c:  c,
		db: mysql.InitDB(c.Mysql),
	}
	return
}

func (r *Repository) Close() {
	log.GetLogger().Info("database successfully closed")
}
