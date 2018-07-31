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
		Name     string `required:"true"`
		Type     string `required:"true" default:"mysql"`
		User     string `required:"true" default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Url      string `required:"true"`
		Port     uint   `default:"3306"`
	}
}{}

var DB *gorm.DB

var DataSource = make(map[string]*gorm.DB, 10)

func GetDataSource(user string, password string, url string) *gorm.DB {
	key := user + "|" + password + "|" + url
	datasource := DataSource[key]
	if datasource == nil {
		connectUrl := fmt.Sprintf("%s:%s@(%s:3306)?charset=utf8&parseTime=True",
			user, password, url)
		datasource, err := gorm.Open("mysql", connectUrl)
		if err != nil {
			panic(err)
		}
		DataSource[key] = datasource
	}
	return datasource
}

func init() {
	configor.Load(&DbConfig, "db.yml")
	connectUrl := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True",
		DbConfig.DB.User, DbConfig.DB.Password, DbConfig.DB.Url, DbConfig.DB.Port, DbConfig.DB.Name)
	db, err := gorm.Open(DbConfig.DB.Type, connectUrl)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(0)
	db.DB().SetConnMaxLifetime(time.Minute * 30)
	db.DB().SetMaxOpenConns(50)
	db.LogMode(true)
	DB = db
}
