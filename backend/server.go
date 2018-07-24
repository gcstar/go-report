package main

import (
	"io"
	"os"
	_ "report/goreport/backend/db"
	"report/goreport/backend/web/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	f, _ := os.Create("log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	r.Use(cors.Default())
	router.AddMetadataController(r)
	r.Run(":7777")
}
