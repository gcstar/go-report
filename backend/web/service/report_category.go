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

func GetAllCategory() *[]Category {
	var categories []Category
	DB.Find(&categories)
	return &categories
}

func EditCategory(category Category) int64 {
	category.GmtModified = time.Now()
	db := DB.Save(&category)
	return db.RowsAffected
}

func AddCategory(category Category) int64 {
	category.GmtCreated = time.Now()
	category.GmtModified = time.Now()
	db := DB.Create(&category)
	return db.RowsAffected
}

func DeleteCategory(category Category) int64 {
	db := DB.Delete(&category)
	return db.RowsAffected
}

func GetCategory(id int) *Category {
	var category Category
	DB.Where("id=?", id).Find(&category)
	return &category
}

func SaveCategoryName(categoryId int, name string) int64 {
	category := Category{ID: categoryId, Name: name}
	db := DB.Model(&category).Select("name").Update(&category)
	return db.RowsAffected
}
