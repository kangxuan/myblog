package v1

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
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

// GetCategory 获取单个分类
func GetCategory(c *gin.Context) {
	appG := util.Gin{C: c}
	var valid validation.Validation

	categoryId := com.StrTo(c.Query("id")).MustInt()
	valid.Min(categoryId, 1, "id").Message("分类ID小于1")

	if valid.HasErrors() {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	category, err := dao.GetCategoryById(categoryId)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, category)
}

func CreateCategory(c *gin.Context) {

}

func UpdateCategory(c *gin.Context) {

}

func DeleteCategory(c *gin.Context) {

}
