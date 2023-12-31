package entity

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestCreateTable(t *testing.T) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"user", "password", "ip:port", "only-id")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal("connect to db fail || reason: ", err.Error())
	}
	db.AutoMigrate(&Segments{})
}
