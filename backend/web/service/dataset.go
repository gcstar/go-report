package service

import (
	"encoding/json"
	. "report/goreport/backend/db"
	"time"
)

type Dataset struct {
	ID           int        `json:"id"`
	UserId       int        `json:"userId"`
	CategoryName string     `json:"categoryName"`
	DatasetName  string     `json:"datasetName"`
	DataJson     string     `json:"dataJson"`
	GmtCreated   *time.Time `json:"gmt_created"`
	GmtModified  *time.Time `json:"gmt_modified"`
}

func (dataset Dataset) TableName() string {
	return "dashboard_dataset"
}

func GetDatasetById(id int) Dataset {
	data_json := make(map[string]interface{})
	var dataset Dataset
	DB.Where("id=?", id).Find(&dataset)
	json.Unmarshal([]byte(dataset.DataJson), &data_json)
	return dataset
}

func AddDataset(data string) int64 {
	var dataset Dataset
	err := json.Unmarshal([]byte(data), &dataset)
	if err != nil {
		panic("parse dataset error!")
	}
	if countExistDatasetName(dataset) == 0 {
		db := DB.Save(dataset)
		return db.RowsAffected
	}
	return 0
}

func countExistDatasetName(dataset Dataset) int64 {
	var count int64
	DB.Where("dateset_name=?", dataset.DatasetName).
		Where("category_name=?", dataset.CategoryName).
		Where("dataset_id<>", dataset.ID).Count(&count)
	return count
}
