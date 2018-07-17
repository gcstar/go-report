package router

import (
	"net/http"
	. "report/goreport/backend/web/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/satori/go.uuid"
)

func AddRoleController(router *gin.Engine) {
	router.POST("/saveRole", saveRole)
	router.GET("/getAllRoles", getAllRoles)
	router.POST("/updateRole", updateRole)
	router.GET("/deleteRole", deleteRole)
	router.GET("/getRoleResList", getRoleResList)
	router.GET("/getUserRoleResList", getUserRoleResList)
}

func saveRole(c *gin.Context) {
	var role Role
	if err := c.ShouldBindWith(&role, binding.JSON); err == nil {
		uuid, _ := uuid.NewV4()
		role.RoleId = uuid.String()
		SaveRole(role)
		c.JSON(http.StatusOK, gin.H{"status": "save role success"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "save role error"})

	}
}

func getAllRoles(c *gin.Context) {
	var roles []Role
	roles = GetAllRoles()
	c.JSON(http.StatusOK, gin.H{"data": roles, "msg": "get roles success", "status": "ok"})
}

func updateRole(c *gin.Context) {
	var role Role
	if err := c.ShouldBindWith(&role, binding.JSON); err == nil {
		UpdateRole(role)
		c.JSON(http.StatusOK, gin.H{"status": "update role success"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "update role error"})

	}
}

func deleteRole(c *gin.Context) {
	roleId := c.DefaultQuery("roleId", "-1")
	DeleteRole(roleId)
}
func getRoleResList(c *gin.Context) {
	roleRes := GetRoleResList()
	if len(roleRes) == 0 {
		c.JSON(http.StatusOK, gin.H{"msg": "no data"})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "get roleRes success", "data": roleRes})

	}
}

func getUserRoleResList(c *gin.Context) {
	userId := c.DefaultQuery("userId", "-1")
	resType := c.DefaultQuery("resType", "-1")
	roleResList := GetUserRoleResList(userId, resType)
	if len(roleResList) == 0 {
		c.JSON(http.StatusOK, gin.H{"msg": "no data"})

	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "get roleRes success", "data": roleResList})

	}
}
