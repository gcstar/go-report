package service

import (
	"fmt"
	"report/goreport/backend/db"
	"strings"
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

	var trimSql string
	if strings.Contains(sqlText, "limit") {
		index := strings.LastIndex(sqlText, "limit")
		trimSql = sqlText[0:index]
	} else {
		trimSql = sqlText
	}

	trimSql = trimSql + " limit 1"

	rows, err := DB.Exec(trimSql).Rows()
	defer rows.Close()
	defer DB.Close()
	if err == nil {
		panic(err)
	}
	columns, err := rows.Columns()
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		panic(err)
	}
	columnList := make([]ReportMetaDataColumn, len(columns))

	for i, column := range columns {
		columnList[i] = ReportMetaDataColumn{
			Ordinal: i,
			Name:    column,
		}
	}
	for i, columnType := range columnTypes {
		metaColumn := columnList[i]
		metaColumn.DataType = columnType.Name()
		length, hasLength := columnType.Length()
		if hasLength {
			metaColumn.Width = length
		} else {
			metaColumn.Width = 0
		}
		columnList[i] = metaColumn
	}

	return &columnList
}

//TODO
func (mysqlQueryer *MysqlQueryer) ParseQueryParamItems(sqlText string) *[]ReportQueryParameter {
	user := mysqlQueryer.dataSource.User
	password := mysqlQueryer.dataSource.Password
	url := mysqlQueryer.dataSource.JdbcUrl
	DB := db.GetDataSource(user, password, url)

	rows, err := DB.Exec(sqlText).Rows()
	if err != nil {
		panic(err)
	}

	fmt.Print(rows)

	return nil
}
