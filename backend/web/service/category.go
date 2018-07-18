package service

import (
	. "report/goreport/backend/db"

	"time"
)

// 报表分类
type Category struct {
	ID          int       `json:"id"`
	ParentId    int       `json:"parentId"`
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	HasChild    *bool     `json:"hasChild"`
	Status      int       `json:"status"`
	Sequence    int       `json:"sequence"`
	Comment     string    `json:"comment"`
	GmtCreated  time.Time `json:"gmtCreated"`
	GmtModified time.Time `json:"gmtModified"`
}

func (category Category) TableName() string {
	return "goreport_category"
}

func GetCategoryList() []Category {
	var categoryList []Category
	DB.Find(&categoryList)
	return categoryList
}

func UpdateCategory(category Category) {
	category.GmtModified = time.Now()
	DB.Save(&category)
}

func AddCategory(category Category) {
	category.GmtCreated = time.Now()
	category.GmtModified = time.Now()
	DB.Create(&category)
}

func DeleteCategory(category Category) {
	DB.Delete(&category)
}

func GetCategory(id int) *Category {
	var category Category
	DB.Where("id=?", id).Find(&category)
	return &category
}
