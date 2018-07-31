package service

import (
	"encoding/json"
	"fmt"
	"reflect"
	"report/goreport/backend/utils"

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
	Width           int64  `json:"width"`
	Decimals        int    `json:"decimals"`
	Type            int    `json:"type"`
	SortType        int    `json:"sortType"`
	IsExtensions    *bool  `json:"isExtensions"`
	IsFootings      *bool  `json:"isFootings"`
	IsPercent       *bool  `json:"isPercent"`
	IsOptional      *bool  `json:"isOptional"`
	IsDisplayInMail *bool  `json:"isDisplayInMail"`
	IsHidden        *bool  `json:"isHidden"`
}

//报表参数选项
type ReportQueryParameter struct {
	Name           string `json:"name"`
	Text           string `json:"text"`
	FormElement    string `json:"formElement"`
	Content        string `json:"content"`
	DefaultValue   string `json:"defaultValue"`
	DefaultText    string `json:"defaultText"`
	DataSource     string `json:"dataSource"`
	DataType       string `json:"dataType"`
	Comment        string `json:"comment"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	IsRequired     *bool  `json:"isRequired"`
	IsAutoComplete *bool  `json:"isAutoComplete"`
}

type ReportParameter struct {
	ID                 string `json:"id"`
	Name               string
	Layout             int
	StatColumnLayout   int    `json:"statColumnLayout"`
	SqlText            string `json:"sqlText"`
	MetaColumns        *[]ReportMetaDataColumn
	EnabledStatColumns *[]string
	IsRowSpan          *bool
}

type ReportQueryParamItem struct {
	Name     string
	Text     string
	Selected *bool
}

type ReportTable struct {
	HtmlText            string `json:"htmlText"`
	SqlText             string `json:"sqlText"`
	MetaDataRowCount    int    `json:"metaDataRowCount"`
	MetaDataColumnCount int    `json:"metaDataColumnCount"`
}

func (report Report) TableName() string {
	return "ezrpt_meta_report"
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

func ListAllReports() *[]Report {
	var reports []Report
	DB.Find(&reports)
	return &reports
}

func SaveQueryParams(reportId int, queryParams string) int64 {
	report := Report{ID: reportId, QueryParams: queryParams}
	db := DB.Model(&report).Select("query_params").Update(&report)
	return db.RowsAffected
}

func GetReportByUid(uid string) *Report {
	var report Report
	DB.Where("uid=?", uid).Find(&report)
	return &report
}

func GetReportSqlText(reportId int) string {
	var report Report
	DB.Where("id=?", reportId).Select("SqlText").Find(&report)
	return report.SqlText
}

func GetReportQueryColumn(uid string) {

}

func GetReportTableData(uid string, params []ReportQueryParameter) map[string]interface{} {

	var res map[string]interface{} = make(map[string]interface{}, 2)

	report := GetReportByUid(uid)
	sqlText := report.SqlText
	var querySql string = sqlText

	metaColumns := parseMetaColumns(report.MetaColumns)

	// if data, err := json.Marshal(metaColumns); err != nil {
	res["metadata"] = metaColumns
	// }

	for _, param := range params {
		var regexpStr = "\\$\\{" + param.Name + "\\}"
		querySql = utils.ReplaceAll(querySql, regexpStr, param.DefaultValue)
	}

	rows, err := DB.Raw(querySql).Rows()
	if err != nil {
		panic("query sql error")
	}

	var results []map[string]interface{}
	cols, err := rows.Columns()

	for rows.Next() {
		var row = make([]interface{}, len(cols))
		var rowp = make([]interface{}, len(cols))

		for i := 0; i < len(cols); i++ {
			rowp[i] = &row[i]
		}
		rows.Scan(rowp...)
		rowMap := make(map[string]interface{})
		for i, col := range cols {
			t := reflect.TypeOf(row[i])
			if t != nil {
				tst := reflect.TypeOf(row[i]).String()
				if tst == "[]uint8" {
					ba := []byte{}
					for _, b := range row[i].([]uint8) {
						ba = append(ba, byte(b))
					}
					rowMap[col] = string(ba)
				} else {
					rowMap[col] = row[i]
				}
			} else {
				rowMap[col] = ""
			}

		}

		results = append(results, rowMap)
	}
	res["data"] = results
	return res

}

func parseOptions(jsonStr string) *ReportOptions {
	var reportOption ReportOptions
	json.Unmarshal([]byte(jsonStr), &reportOption)
	return &reportOption
}

func parseQueryParams(jsonStr string) []ReportQueryParameter {
	var queryParameter []ReportQueryParameter
	json.Unmarshal([]byte(jsonStr), &queryParameter)
	return queryParameter
}

func parseMetaColumns(jsonStr string) []ReportMetaDataColumn {
	var metaDataColumn []ReportMetaDataColumn
	json.Unmarshal([]byte(jsonStr), &metaDataColumn)
	return metaDataColumn
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
