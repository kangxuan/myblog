package v1

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"myblog/dao"
	"myblog/models"
	"myblog/pkg/e"
	"myblog/pkg/util"
	"net/http"
	"reflect"
)

func GetCategoryList(c *gin.Context) {
	appG := util.Gin{
		C: c,
	}
	maps := make(map[string]string)

	if arg := c.Query("category_name"); arg != "" {
		maps["category_name"] = arg
	}
	maps["page"] = c.DefaultQuery("page", "1")
	maps["page_size"] = c.DefaultQuery("page_size", "10")

	categoryList, err := dao.GetCategoryList(maps)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, "获取分类列表失败", nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, "", map[string]interface{}{
		"list": categoryList,
		"page": dao.GetCategoryPage(maps),
	})
}

// GetCategory 获取单个分类
func GetCategory(c *gin.Context) {
	appG := util.Gin{C: c}
	var valid validation.Validation

	categoryId := com.StrTo(c.Param("id")).MustInt()
	fmt.Println("category_id:", categoryId)
	valid.Min(categoryId, 0, "id").Message("分类ID小于1")

	hasError, msg := util.GetValidationMessage(valid)
	if hasError {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, msg, nil)
		return
	}

	category, err := dao.GetCategoryById(categoryId)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, "", nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, "", category)
}

// CreateCategory 创建分类
func CreateCategory(c *gin.Context) {
	appG := util.Gin{C: c}
	var valid validation.Validation

	var categoryParams models.Category
	err := c.BindJSON(&categoryParams)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, "4234234", nil)
		return
	}

	fmt.Println(reflect.TypeOf(categoryParams.CategoryType))
	valid.Required(categoryParams.CategoryName, "category_name").Message("分类名称不能为空")
	valid.MaxSize(categoryParams.CategoryName, 20, "category_name").Message("分类名称长度不能超过20个字")
	valid.Required(categoryParams.CategoryType, "category_type").Message("分类类型不能为空")
	valid.Range(categoryParams.CategoryType, 1, 3, "category_type").Message("分类类型错误")
	valid.Min(categoryParams.ParentId, 0, "parent_id").Message("父级ID错误")

	hasError, msg := util.GetValidationMessage(valid)
	if hasError {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, msg, nil)
		return
	}

	lastCategoryId, err := dao.InsertCategory(categoryParams)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, "", map[string]int{
		"category_id": lastCategoryId,
	})
}

func UpdateCategory(c *gin.Context) {

}

func DeleteCategory(c *gin.Context) {

}
