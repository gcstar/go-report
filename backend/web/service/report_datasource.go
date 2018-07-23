package service

import (
	"encoding/json"
	. "report/goreport/backend/db"

	"time"
)

// 数据源
type Datasource struct {
	ID           int       `json:"id"`
	Uid          string    `json:"uid"`
	Name         string    `json:"name"`
	DriverClass  string    `json:"driverClass"`
	JdbcUrl      string    `json:"jdbcUrl"`
	User         string    `json:"user"`
	Password     string    `json:"password"`
	QueryerClass string    `json:"queryerClass"`
	PoolClass    string    `json:"poolClass"`
	Options      string    `json:"options"`
	Comment      string    `json:"comment"`
	GmtCreated   time.Time `json:"gmtCreated"`
	GmtModified  time.Time `json:"gmtModified"`
}

//数据连接池配置选项
type ConnectionPoolOptions struct {
	InitialSize                   int   `json:"initialSize"`
	MaxActive                     int   `json:"maxActive"`
	MinIdle                       int   `json:"minIdle"`
	MaxWait                       int   `json:"maxWait"`
	TimeBetweenEvictionRunsMillis int   `json:"timeBetweenEvictionRunsMillis"`
	MinEvictableIdleTimeMillis    int   `json:"minEvictableIdleTimeMillis"`
	TestWhileIdle                 *bool `json:"testWhileIdle"`
	TestOnBorrow                  *bool `json:"testOnBorrow"`
	TestOnReturn                  *bool `json:"testOnReturn"`
	MaxOpenPreparedStatements     int   `json:"maxOpenPreparedStatements"`
	RemoveAbandoned               *bool `json:"removeAbandoned"`
	RemoveAbandonedTimeout        int   `json:"removeAbandonedTimeout"`
	LogAbandoned                  *bool `json:"logAbandoned"`
}

type ReportDatasource struct {
	Uid          string                `json:"uid"`
	QueryerClass string                `json:"queryerClass"`
	DbPoolClass  string                `json:"dbPoolClass"`
	DriverClass  string                `json:"driverClass"`
	JdbcUrl      string                `json:"jdbcUrl"`
	User         string                `json:"user"`
	Password     string                `json:"password"`
	Options      ConnectionPoolOptions `json:"options"`
}

func (datasource Datasource) TableName() string {
	return "ezrpt_meta_datasource"
}

func AddDatasource(datasource Datasource) int64 {
	datasource.GmtCreated = time.Now()
	datasource.GmtModified = time.Now()
	db := DB.Create(&datasource)
	return db.RowsAffected
}

func GetAllDatasource() *[]Datasource {
	var datasourceList []Datasource
	DB.Find(&datasourceList)
	return &datasourceList
}

func EditDatasource(datasource Datasource) int64 {
	datasource.GmtModified = time.Now()
	db := DB.Save(&datasource)
	return db.RowsAffected
}

func DeleteDatasource(id int) int64 {
	var datasource Datasource
	db := DB.Where("id=?", id).Find(&datasource).Delete(&datasource)
	return db.RowsAffected
}

func GetReportDatasource(dsId int) *ReportDatasource {
	var datasource Datasource
	DB.Where("id=?", dsId).Find(&datasource)
	var options ConnectionPoolOptions
	json.Unmarshal([]byte(datasource.Options), &options)
	reportDatasource := ReportDatasource{
		Uid:          datasource.Uid,
		QueryerClass: datasource.QueryerClass,
		DbPoolClass:  datasource.PoolClass,
		DriverClass:  datasource.DriverClass,
		JdbcUrl:      datasource.JdbcUrl,
		User:         datasource.User,
		Password:     datasource.Password,
		Options:      options,
	}
	return &reportDatasource
}

//TODO
func TestConnection(datasource Datasource) bool {
	// url := datasource.JdbcUrl
	// user := datasource.User
	// password := datasource.Password
	return true
}
