package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DbConfig = struct {
	DB struct {
		Name     string
		Type     string `default:"mysql"`
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Url      string `required:"true"`
		Port     uint   `default:"3306"`
	}
}{}

var DB *gorm.DB

func init() {
	configor.Load(&DbConfig, "db.yml")
	connectUrl := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True",
		DbConfig.DB.User, DbConfig.DB.Password, DbConfig.DB.Url, DbConfig.DB.Port, DbConfig.DB.Name)
	db, err := gorm.Open(DbConfig.DB.Type, connectUrl)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(0)
	db.DB().SetConnMaxLifetime(time.Minute * 5)
	db.DB().SetMaxOpenConns(50)
	db.LogMode(true)
	DB = db
}
