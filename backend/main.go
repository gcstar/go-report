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
	r.LoadHTMLGlob("../frontend/views/**/*")
	AddRoleController(r)
	AddViewsController(r)
	r.Run(":7777")
}
