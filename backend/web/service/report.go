package service

import (
	"fmt"
	. "report/goreport/backend/db"
	"time"
)

type ColumnType int
type ColumnSortType int

const (
	LAYOUT ColumnType = iota + 1
	DIMENSION
	STATISTICAL
	COMPUTED
)

const (
	DEFAULT ColumnSortType = iota
	DIGIT_ASCENDING
	DIGIT_DESCENDING
	CHAR_ASCENDING
	CHAR_DESCENDING
)

// 报表实体
type Report struct {
	ID           int       `json:"id"`
	Uid          string    `json:"uid"`
	CategoryId   int       `json:"categoryId"`
	DsId         int       `json:"dsId"`
	Name         string    `json:"name"`
	SqlText      string    `json:"sqlText"`
	MetaColumns  string    `json:"metaColumns"`
	QueryParams  string    `json:"queryParams"`
	Options      string    `json:"options"`
	Status       int       `json:"status"`
	Sequence     int       `json:"sequence"`
	Comment      string    `json:"comment"`
	CreateUser   string    `json:"createUser"`
	GmtCreated   time.Time `json:"gmtCreated"`
	GmtModified  time.Time `json:"gmtModified"`
	CategoryName string    `gorm:"-" json:"categoryName"`
	DsName       string    `gorm:"-" json:"dsName"`
}

// 报表选项
type ReportOptions struct {
	DataRange        int `json:"dataRange"`
	Layout           int `json:"layout"`
	StatColumnLayout int `json:"statColumnLayout"`
}

//报表元数据列类
type ReportMetaDataColumn struct {
	Ordinal         int    `json:"ordinal"`
	Name            string `json:"name"`
	Text            string `json:"text"`
	DataType        string `json:"dataType"`
	Expression      string `json:"expression"`
	Format          string `json:"format"`
	Comment         string `json:"comment"`
	Width           int    `json:"width"`
	Decimals        int    `json:"decimals"`
	Type            int    `json:"type"`
	SortType        int    `json:"sortType"`
	IsExtensions    *bool  `json:"isExtensions"`
	IsFootings      *bool  `json:"isFootings"`
	IsPercent       *bool  `json:"isPercent"`
	IsOptional      *bool  `json:"isOptional"`
	IsDisplayInMail *bool  `json:"isDisplayInMail"`
	UsHidden        *bool  `json:"isHidden"`
}

func (report Report) TableName() string {
	return "goreport_report"
}

func ListReportByCategoryId(categoryId int) *[]Report {
	var reports []Report
	DB.Where("category_id=?", categoryId).Find(&reports)
	return &reports
}

func FindReport(fieldName string, keyWord string) *[]Report {
	var reports []Report
	DB.Where(fieldName+" like ?", "%"+keyWord+"%").Find(&reports)
	return &reports
}

func EditReport(report Report) int64 {
	db := DB.Save(&report)
	return db.RowsAffected
}

func AddReport(report Report) int64 {
	report.GmtCreated = time.Now()
	report.GmtModified = time.Now()
	db := DB.Create(&report)
	return db.RowsAffected
}
func DeleteReport(id int) int64 {
	var report Report
	db := DB.Where("id=?", id).Find(&report).Delete(&report)
	return db.RowsAffected
}

func SaveQueryParams(reportId int, queryParams string) int64 {
	report := Report{ID: reportId, QueryParams: queryParams}
	db := DB.Model(&report).Select("query_params").Update(&report)
	return db.RowsAffected
}

func GetReportByUid(uid int) *[]Report {
	var reports []Report
	DB.Where("uid=?", uid).Find(&reports)
	return &reports
}

func GetReportSqlText(reportId int) string {
	var report Report
	DB.Where("id=?", reportId).Select("SqlText").Find(&report)
	return report.SqlText
}

func GetMetaDataColumns(dsId int, sqlText string) {

}

func ExplainSqlText(dsId int, sqlText string) {

}

func ExecuteSqlText(dsId int, sqlText string) {
	reportDatasource := GetReportDatasource(dsId)
	fmt.Println(reportDatasource)

}

func ExecuteQueryParamSqlText(dsId int, sqlText string) {

}
