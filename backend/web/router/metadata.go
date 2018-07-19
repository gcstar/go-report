package router

import (
	"net/http"
	. "report/goreport/backend/web/model"
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
		metadata.GET("/report/findReport", findReport)
		metadata.POST("/report/editReport", editReport)
	}
}

func listReportByCategoryId(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.DefaultQuery("categoryId", "-1"))
	reports := ListReportByCategoryId(categoryId)
	if reports == nil {
		c.JSON(200, NoData())
	} else {
		c.JSON(200, OK("get categories success", reports))
	}
}

func getCategoryList(c *gin.Context) {
	categoryList := GetCategoryList()
	if len(categoryList) != 0 {
		c.JSON(http.StatusOK, OK("get categoryList success", categoryList))
	} else {
		c.JSON(http.StatusOK, NoData())
	}
}

func updateCategory(c *gin.Context) {
	var category Category
	if err := c.BindJSON(&category); err == nil {
		UpdateCategory(category)
		c.JSON(http.StatusOK, OK("update category success", nil))
	} else {
		c.JSON(http.StatusBadRequest, ParameterError())
	}
}

func addCategory(c *gin.Context) {
	var category Category
	if err := c.BindJSON(&category); err == nil {
		AddCategory(category)
		c.JSON(http.StatusOK, OK("add category success", nil))
	} else {
		c.JSON(http.StatusBadRequest, ParameterError())
	}
}

func deleteCategory(c *gin.Context) {
	var category Category
	if err := c.BindJSON(&category); err == nil {
		DeleteCategory(category)
		c.JSON(http.StatusOK, OK("delete category success", nil))
	} else {
		c.JSON(http.StatusBadRequest, ParameterError())
	}
}

func getCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category := GetCategory(id)
	if category != nil {
		c.JSON(200, OK("get category success", category))
	} else {
		c.JSON(200, NoData())
	}
}

func findReport(c *gin.Context) {
	fieldName := c.DefaultQuery("fieldName", "name")
	keyWord := c.DefaultQuery("keyWord", "report")
	reports := FindReport(fieldName, keyWord)
	if reports == nil {
		c.JSON(200, NoData())
	} else {
		c.JSON(200, OK("find reports success", reports))
	}
}

func editReport(c *gin.Context) {
	var report Report
	if err := c.BindJSON(&report); err == nil {
		row := EditReport(report)
		if row != 0 {
			c.JSON(200, OK("update report success", nil))
		} else {
			c.JSON(200, OK("no report updated", nil))
		}
	} else {
		c.JSON(http.StatusBadRequest, ParameterError())
	}
}
