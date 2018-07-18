package service

import (
	. "report/goreport/backend/db"
)

type Role struct {
	RoleId   string `json:"roleId"`
	RoleName string `json:"roleName" binding:"required"`
	UserId   string `json:"userId" binding:"required"`
}

type RoleRes struct {
	RoleResId  uint32 `json:"roleResId"`
	RoleId     string `json:"roleId"`
	ResId      uint32 `json:"resId"`
	ResType    string `json:"resType"`
	Permission string `json:"permission"`
}

func (role Role) TableName() string {
	return "report_role"
}

func (roleRes RoleRes) TableName() string {
	return "report_role_res"
}

func SaveRole(role Role) {
	DB.Create(&role)
}

func GetAllRoles() []Role {
	var roles []Role
	DB.Find(&roles)
	return roles
}

func GetRole(roleId string) Role {
	var role Role
	DB.Where("role_id=?", roleId).Find(&role)
	return role
}

func UpdateRole(role Role) {
	var oldRole Role
	DB.Where("role_id=?", role.RoleId).Find(&oldRole).Update(&role)
}

func DeleteRole(roleId string) {
	var role Role
	DB.Where("role_id=?", roleId).Find(&role).Delete(&role)
}

func GetRoleResList() []RoleRes {
	var roleRes []RoleRes
	DB.Find(&roleRes)
	return roleRes
}

func GetUserRoleResList(userId string, resType string) []RoleRes {
	var roleResList []RoleRes
	DB.Table("dashboard_user_role").
		Joins("left join dashboard_role_res on dashboard_user_role.role_id=dashboard_role_res.role_id").
		Where("dashboard_user_role.user_id=?", userId).
		Where("dashboard_role_res.res_type=?", resType).
		Find(&roleResList)
	return roleResList
}

func SaveRoleRes(roleResList []RoleRes) {
	for _, roleRes := range roleResList {
		DB.Table("report_role_res").Create(&roleRes)
	}
}

func DeleteRoleRes(roleId string) {
	DB.Where("role_id=?", roleId).Delete(RoleRes{})
}
