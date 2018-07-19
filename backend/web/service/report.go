package service

import "time"
import . "report/goreport/backend/db"

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
