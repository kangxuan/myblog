package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"myblog/dao"
	"myblog/pkg/e"
	"myblog/pkg/util"
	"myblog/settings"
	"net/http"
)

func GetCategoryList(c *gin.Context) {
	appG := util.Gin{
		C: c,
	}
	maps := make(map[string]interface{})

	if arg := c.Query("category_name"); arg != "" {
		maps["category_name"] = arg
	}

	maps["page"] = util.GetPage(c)
	maps["page_size"] = settings.ServerConf.PageSize
	fmt.Println(maps)

	categoryList, err := dao.GetCategoryList(maps)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusOK, e.ERROR, "获取分类失败")
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, categoryList)
}

func GetCategory(c *gin.Context) {

}

func CreateCategory(c *gin.Context) {

}

func UpdateCategory(c *gin.Context) {

}

func DeleteCategory(c *gin.Context) {

}
