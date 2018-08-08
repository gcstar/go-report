package service

import (
	"encoding/json"

	"strconv"
)

type ValueConfig struct {
	ColumnName string `json:"columnName"`
	AggType    string `json:"aggType"`
}

type AggConfig struct {
	Rows    *[]DimensionConfig `json:"rows"`
	Columns *[]DimensionConfig `json:"columns"`
	Filters *[]DimensionConfig `json:"filters"`
	Values  *[]ValueConfig     `json:"values"`
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
	ColumnName string    `json:"columnName"`
	FilterType string    `json:"filterType"`
	Values     *[]string `json:"values"`
}

type Aggregatable interface {
	GetColumn() []string
	QueryDataAggData(config AggConfig)
	QueryDimVals(columnName string, config AggConfig) []string
}

type SimpleDataSet struct {
	DatasourceId int    `json:"datasourceId"`
	Query        string `json:"query"`
	Interval     int    `json:"interval"`
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
