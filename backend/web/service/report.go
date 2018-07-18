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
	CategoryName string    `json:"categoryName"`
	DsName       string    `json:"dsName"`
}

func (report Report) TableName() string {
	return "goreport_report"
}

func ListReportByCategoryId(categoryId int, page int, row int) *[]Report {
	var reports []Report
	offset := row * (page - 1)
	DB.Where("category_id=?", categoryId).Offset(offset).Limit(row).Find(&reports)
	return &reports
}
