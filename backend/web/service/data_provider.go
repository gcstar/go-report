package service

import (
	"encoding/json"
	"report/goreport/backend/utils"
	"strings"

	// . "report/goreport/backend/db"
	"strconv"
)

type Aggregatable interface {
	GetColumn() []string
	QueryDataAggData(config AggConfig)
	QueryDimVals(columnName string, config AggConfig) []string
}

type ValueConfig struct {
	ColumnName string `json:"column"`
	AggType    string `json:"aggType"`
}

type AggConfig struct {
	Rows    []DimensionConfig `json:"rows"`
	Columns []DimensionConfig `json:"columns"`
	Filters []DimensionConfig `json:"filters"`
	Values  []ValueConfig     `json:"values"`
}

type AggregateResult struct {
	Data       [][]string
	ColumnList []ColumnIndex
}
type ColumnIndex struct {
	Index   int
	AggType string
	Name    string
}

type DimensionConfig struct {
	ColumnName string   `json:"columnName"`
	FilterType string   `json:"filterType"`
	Values     []string `json:"values"`
}

type SimpleDataSet struct {
	DatasourceId int    `json:"datasourceId"`
	Query        string `json:"query"`
	Interval     int    `json:"interval"`
}

func TestQueryAggData(configStr string, sql string) string {
	config := GetAggConfig(configStr)
	return getQueryAggDataSql(config, sql)
}

func GetAggConfig(configStr string) AggConfig {
	var config AggConfig
	json.Unmarshal([]byte(configStr), &config)
	return config
}

func GetDataProvider(datasourceId int, query map[string]string, datasetId int) {
	if datasetId != 0 {
		dataset := GetDatasetById(datasetId)
		var simpleDataset SimpleDataSet
		simpleDataset.DatasourceId = dataset.ID
		jsonMap := make(map[string]interface{})
		json.Unmarshal([]byte(dataset.DataJson), &jsonMap)
		if query, ok := jsonMap["query"]; ok {
			if sql, ok := query.(map[string]interface{})["sql"]; ok {
				simpleDataset.Query = sql.(string)
			}
		}
		if interval, ok := jsonMap["interval"]; ok {
			value, err := strconv.Atoi(interval.(string))
			if err == nil {
				simpleDataset.Interval = value
			}
		}

		datasetId = simpleDataset.DatasourceId
	}
}

func queryAggData(config AggConfig) {

}

//TODO 自定义参数解析
func getQueryAggDataSql(config AggConfig, sql string) string {
	columns := config.Columns
	rows := config.Rows
	var filters = make([]DimensionConfig, 0)
	filters = append(append(filters, columns...), rows...)
	values := config.Values

	var dimColsStr = ""
	var aggColsStr = ""
	var whereStr = " WHERE 1=1 AND "
	var groupByStr = ""
	var dimArray = make([]string, 0)
	for _, column := range columns {
		dimArray = append(dimArray, column.ColumnName)
	}
	for _, row := range rows {
		dimArray = append(dimArray, row.ColumnName)
	}
	dimColsStr = strings.Join(dimArray, " , ")

	var aggArray = make([]string, 0)

	for _, value := range values {
		var aggExp string = value.ColumnName
		switch value.AggType {
		case "sum":
			aggArray = append(aggArray, "SUM("+aggExp+")")
		case "count":
			aggArray = append(aggArray, "COUNT("+aggExp+")")
		case "max":
			aggArray = append(aggArray, "MAX("+aggExp+")")
		case "min":
			aggArray = append(aggArray, "min("+aggExp+")")
		case "distinct":
			aggArray = append(aggArray, "SUM("+aggExp+")")
		case "avg":
			aggArray = append(aggArray, "AVG("+aggExp+")")
		default:
		}
	}
	aggColsStr = strings.Join(aggArray, " , ")

	var whereArray []string = make([]string, 0)
	for _, where := range filters {
		if len(where.Values) != 0 {
			str := ""
			switch where.FilterType {
			case "=", "eq":
				strs := strings.Join(where.Values, "','")
				str = where.ColumnName + " IN ('" + strs + "')"
			case "!=", "ne":
				strs := strings.Join(where.Values, "','")
				str = where.ColumnName + " NOT IN ('" + strs + "')"
			case ">":
				str = where.ColumnName + " > '" + where.Values[0] + "')"
			case "<":
				str = where.ColumnName + " <= '" + where.Values[0] + "')"
			case ">=":
				str = where.ColumnName + " >= '" + where.Values[0] + "')"
			case "<=":
				str = where.ColumnName + " <= '" + where.Values[0] + "')"
			case "(a,b]":
				str = where.ColumnName + " > '" + where.Values[0] + "' AND " + where.ColumnName + " <= '" + where.Values[1] + "'"
			case "[a,b)":
				str = where.ColumnName + " >= '" + where.Values[0] + "' AND " + where.ColumnName + " < '" + where.Values[1] + "'"
			case "(a,b)":
				str = where.ColumnName + " > '" + where.Values[0] + "' AND " + where.ColumnName + " < '" + where.Values[1] + "'"
			case "[a,b]":
				str = where.ColumnName + " >= '" + where.Values[0] + "' AND " + where.ColumnName + " <= '" + where.Values[1] + "'"
			default:
			}
			whereArray = append(whereArray, str)
		}
	}

	if len(whereArray) == 0 {
		whereStr = ""
	} else {
		whereStr += strings.Join(whereArray, " AND ")
	}
	if dimColsStr != "" {
		groupByStr = " GROUP BY " + dimColsStr
	}
	var selectColsStr string = ""
	if dimColsStr != "" {
		selectColsStr = dimColsStr + " , " + aggColsStr
	} else {
		selectColsStr = aggColsStr
	}
	querySql := " SELECT " + selectColsStr + " FROM ( " + sql + "  ) as __view__  " + whereStr + " " + groupByStr
	return utils.ReplaceAll(querySql, "\r|\n", " ")
}
