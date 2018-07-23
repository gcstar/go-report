package service

import (
	"report/goreport/backend/db"
)

type MysqlQueryer struct {
	dataSource ReportDatasource
	parameter  ReportQueryParamItem
}
type Queryer interface {
	ParseMetaDataColumns(sqlText string) *[]ReportMetaDataColumn
	ParseQueryParamItems(sqlText string) *[]ReportQueryParamItem
	GetMetaDataColumns() *[]ReportMetaDataColumn
}

//TODO
func (mysqlQueryer *MysqlQueryer) ParseMetaDataColumns(sqlText string) *[]ReportMetaDataColumn {
	user := mysqlQueryer.dataSource.User
	password := mysqlQueryer.dataSource.Password
	url := mysqlQueryer.dataSource.JdbcUrl
	DB := db.GetDataSource(user, password, url)
	rows, err := DB.Exec(sqlText).Rows()
	defer rows.Close()
	if err == nil {
		panic(err)
	}

	for rows.Next() {

	}

	return nil
}
