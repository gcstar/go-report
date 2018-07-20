package service

import "time"
import . "report/goreport/backend/db"

type ReportHistory struct {
	ID          int       `json:"id"`
	ReportId    int       `json:"reportId"`
	Uid         string    `json:"uid"`
	CategoryId  int       `json:"categoryId"`
	DsId        int       `json:"dsId"`
	Name        string    `json:"name"`
	SqlText     string    `json:"sqlText"`
	MetaColumns string    `json:"metaColumns"`
	QueryParams string    `json:"queryParams"`
	Options     string    `json:"options"`
	Status      int       `json:"status"`
	Sequence    int       `json:"sequence"`
	Comment     string    `json:"comment"`
	Author      string    `json:"author"`
	GmtCreated  time.Time `json:"gmtCreated"`
	GmtModified time.Time `json:"gmtModified"`
}

func (history ReportHistory) TableName() string {
	return "goreport_history"
}

func ListReportHistory(reportId int) *[]ReportHistory {
	var reportHistories []ReportHistory
	DB.Where("report_id=?", reportId).Find(&reportHistories)
	return &reportHistories
}
