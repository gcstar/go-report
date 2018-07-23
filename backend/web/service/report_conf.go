package service

import (
	. "report/goreport/backend/db"
	"time"
)

// 数据配置
type Conf struct {
	ID          int       `json:"id"`
	ParentId    int       `json:"parentId"`
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	Key         string    `json:"key"`
	Value       string    `json:"value"`
	Sequence    int       `json:"sequence"`
	Comment     string    `json:"comment"`
	GmtCreated  time.Time `json:"gmtCreated"`
	GmtModified time.Time `json:"gmtModified"`
	HasChild    *bool     `gorm:"-" json:"hasChild"`
}

func (conf Conf) TableName() string {
	return "ezrpt_meta_conf"
}
func FindConf(fieldName string, keyWord string) *[]Conf {
	var confList []Conf
	DB.Where(fieldName+" like ?", "%"+keyWord+"%").Find(&confList)
	return &confList
}

func AddConf(conf Conf) int64 {
	conf.GmtCreated = time.Now()
	conf.GmtModified = time.Now()
	db := DB.Create(&conf)
	return db.RowsAffected
}

func EditConf(conf Conf) int64 {
	conf.GmtModified = time.Now()
	db := DB.Save(&conf)
	return db.RowsAffected
}

func GetAllConf() *[]Conf {
	var confs []Conf
	DB.Find(&confs)
	return &confs
}

func GetConfItem(pid int) *[]Conf {
	var confs []Conf
	DB.Where("pid = ?", pid).Find(&confs)
	return &confs
}
