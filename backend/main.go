package main

import (
	"io"
	"os"
	_ "report/goreport/backend/db"

	. "report/goreport/backend/web/router"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	f, _ := os.Create("log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	AddRoleController(r)
	r.Run(":7777")

	// db, err := gorm.Open("mysql", "growthdb:devhz123456@(172.17.60.131:3306)/bi?charset=utf8&parseTime=True")
	// defer db.Close()
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// uuid, _ := uuid.NewV4()
	// role := Role{RoleId: uuid.String(), RoleName: "test", UserId: "111"}
	// fmt.Println(role)
	// SaveRole(role, db)

	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// db.LogMode(true)
	// var user User
	// db.Where("account=?", "guchao").First(&user)
	// fmt.Println(user.Account)

}
