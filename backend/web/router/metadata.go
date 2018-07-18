package router

import (
	"net/http"
	. "report/goreport/backend/web/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddMetadataController(router *gin.Engine) {
	metadata := router.Group("/metadata")
	{
		metadata.PUT("/category", addCategory)
		metadata.POST("/category", updateCategory)
		metadata.DELETE("/category", deleteCategory)
		metadata.GET("/category/:id", getCategory)
		metadata.GET("/categorys", getCategoryList)

		metadata.GET("/report/list", listReportByCategoryId)

	}
}

func listReportByCategoryId(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	rows, _ := strconv.Atoi(c.DefaultQuery("rows", "10"))
	categoryId, _ := strconv.Atoi(c.DefaultQuery("categoryId", "-1"))
	reports := ListReportByCategoryId(categoryId, page, rows)
	if reports == nil {
		c.JSON(200, gin.H{"msg": "get reports fail"})
	} else {
		c.JSON(200, gin.H{"msg": "get reports success", "data": reports})
	}
}

func getCategoryList(c *gin.Context) {
	categoryList := GetCategoryList()
	if len(categoryList) != 0 {
		c.JSON(http.StatusOK, gin.H{"msg": "get categoryList success", "data": categoryList})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "categoryList empty"})
	}
}

func updateCategory(c *gin.Context) {
	var category Category
	if err := c.BindJSON(&category); err == nil {
		UpdateCategory(category)
		c.JSON(http.StatusOK, gin.H{"msg": "update category success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bad request"})
	}
}

func addCategory(c *gin.Context) {
	var category Category
	if err := c.BindJSON(&category); err == nil {
		AddCategory(category)
		c.JSON(http.StatusOK, gin.H{"msg": "add category success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bad request"})
	}
}

func deleteCategory(c *gin.Context) {
	var category Category
	if err := c.BindJSON(&category); err == nil {
		DeleteCategory(category)
		c.JSON(http.StatusOK, gin.H{"msg": "de category success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bad request"})
	}
}

func getCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category := GetCategory(id)
	if category != nil {
		c.JSON(200, gin.H{"msg": "get category success", "data": category})
	} else {
		c.JSON(200, gin.H{"msg": "get category fail"})

	}

}
