package db

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:123456@(localhost:3306)/bi?charset=utf8&parseTime=True")
	if err != nil {
		panic("init database error")
	}
	db.DB().SetMaxIdleConns(0)
	db.DB().SetConnMaxLifetime(time.Minute * 5)
	db.DB().SetMaxOpenConns(50)
	db.LogMode(true)
	DB = db

}
