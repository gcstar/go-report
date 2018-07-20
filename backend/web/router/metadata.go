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
		metadata.POST("/category", editCategory)
		metadata.DELETE("/category", deleteCategory)
		metadata.GET("/category/:id", getCategory)
		metadata.GET("/categorys", getCategoryList)
		metadata.POST("/category/changename", saveCategoryName)
		metadata.GET("/report/list", listReportByCategoryId)
		metadata.GET("/report/findReport", findReport)
		metadata.POST("/report/editReport", editReport)
		metadata.GET("/datasource", getReportDatasource)
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
	categoryList := GetAllCategory()
	if categoryList != nil {
		c.JSON(http.StatusOK, OK("get categoryList success", categoryList))
	} else {
		c.JSON(http.StatusOK, NoData())
	}
}

func editCategory(c *gin.Context) {
	var category Category
	if err := c.BindJSON(&category); err == nil {
		row := EditCategory(category)
		if row > 0 {
			c.JSON(http.StatusOK, OK("update category success", nil))
		} else {
			c.JSON(http.StatusOK, OK("no category to update", nil))
		}
	} else {
		c.JSON(http.StatusBadRequest, ParameterError())
	}
}

func addCategory(c *gin.Context) {
	var category Category
	if err := c.BindJSON(&category); err == nil {
		row := AddCategory(category)
		if row > 0 {
			c.JSON(http.StatusOK, OK("add category success", nil))
		} else {
			c.JSON(http.StatusOK, OK("add category failed", nil))
		}
	} else {
		c.JSON(http.StatusBadRequest, ParameterError())
	}
}

func deleteCategory(c *gin.Context) {
	var category Category
	if err := c.BindJSON(&category); err == nil {
		row := DeleteCategory(category)
		if row > 0 {
			c.JSON(http.StatusOK, OK("delete category success", nil))
		} else {
			c.JSON(http.StatusOK, OK("delete category fail", nil))
		}
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

func saveCategoryName(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.DefaultQuery("id", "-1"))
	name := c.DefaultQuery("name", "default")
	row := SaveCategoryName(categoryId, name)
	if row > 0 {
		c.JSON(http.StatusOK, OK("update category name success", nil))
	} else {
		c.JSON(http.StatusOK, OK("update category name fail", nil))

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

func addReport(c *gin.Context) {
	var report Report
	if err := c.BindJSON(&report); err == nil {
		row := AddReport(report)
		if row != 0 {
			c.JSON(200, OK("add report success", nil))
		} else {
			c.JSON(200, OK("add report failed", nil))
		}
	} else {
		c.JSON(http.StatusBadRequest, ParameterError())
	}
}

func deleteReport(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "-1"))
	row := DeleteReport(id)
	if row == 0 {
		c.JSON(200, OK("no report to delete", nil))
	} else {
		c.JSON(200, OK("delete report success", nil))
	}
}

func getReportDatasource(c *gin.Context) {
	dsId, _ := strconv.Atoi(c.DefaultQuery("id", "-1"))
	reportDatasource := GetReportDatasource(dsId)
	// if reportDatasource != nil {
	c.JSON(200, OK("get report datasource success", reportDatasource))
	// }
}