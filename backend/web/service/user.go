package service

import (
	"time"
)

type User struct {
	ID          uint
	Roles       string
	Account     string
	Password    string
	Salt        string
	Name        string
	Email       string
	Telephone   string
	Status      bool
	Comment     string
	GmtCreated  time.Time
	GmtModified time.Time
	Department  uint
}

func (User) TableName() string {
	return "ezrpt_member_user"
}
