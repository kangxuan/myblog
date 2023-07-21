package v1

import (
	"errors"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"myblog/dao"
	"myblog/models"
	"myblog/pkg/e"
	"myblog/pkg/util"
	"net/http"
	"strconv"
)

// GetArticleList 获取文章列表
func GetArticleList(c *gin.Context) {
	searchParams := new(models.ArticleSearchColumns)
	searchParams.Title = c.DefaultQuery("title", "")
	searchParams.CategoryId, _ = strconv.Atoi(c.DefaultQuery("category_id", "0"))
	searchParams.TagId, _ = strconv.Atoi(c.DefaultQuery("tag_id", "0"))
	searchParams.Page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	searchParams.PageSize, _ = strconv.Atoi(c.DefaultQuery("page_size", "10"))

	articleList, err := dao.GetArticleList(searchParams)
	if err != nil {
		util.Response(c, http.StatusOK, e.ERROR, err.Error(), nil)
		return
	}
	util.Response(c, http.StatusOK, e.SUCCESS, "获取文章列表成功", map[string]interface{}{
		"list": articleList,
		"page": dao.GetArticlePage(searchParams),
	})
}

// GetArticle 获取文章详情
func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := dao.GetArticle(id)
	if err != nil {
		util.Response(c, http.StatusOK, e.ERROR, err.Error(), nil)
		return
	}
	util.Response(c, http.StatusOK, e.SUCCESS, "获取文章详情成功", article)
}

// CreateArticle 创建文章
func CreateArticle(c *gin.Context) {
	var params *models.ArticleSaveColumns
	err := c.BindJSON(&params)
	if err != nil {
		util.Response(c, http.StatusOK, e.ERROR, "参数解析错误", nil)
		return
	}

	err = checkArticleParams(params)
	if err != nil {
		util.Response(c, http.StatusOK, e.INVALID_PARAMS, err.Error(), nil)
		return
	}

	articleId, err := dao.CreateArticle(params)
	if err != nil {
		util.Response(c, http.StatusOK, e.ERROR, err.Error(), nil)
		return
	}

	util.Response(c, http.StatusOK, e.SUCCESS, "创建文章成功", map[string]interface{}{
		"article_id": articleId,
	})
}

// UpdateArticle 修改文章
func UpdateArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var params *models.ArticleSaveColumns
	err := c.BindJSON(&params)
	if err != nil {
		util.Response(c, http.StatusOK, e.ERROR, "参数解析错误", nil)
		return
	}

	err = checkArticleParams(params)
	if err != nil {
		util.Response(c, http.StatusOK, e.INVALID_PARAMS, err.Error(), nil)
		return
	}

	err = dao.UpdateArticle(id, params)
	if err != nil {
		util.Response(c, http.StatusOK, e.ERROR, err.Error(), nil)
		return
	}

	util.Response(c, http.StatusOK, e.SUCCESS, "修改文章成功", nil)
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := dao.DeleteArticle(id)
	if err != nil {
		util.Response(c, http.StatusOK, e.ERROR, err.Error(), nil)
		return
	}
	util.Response(c, http.StatusOK, e.SUCCESS, "删除文章成功", nil)
}

// checkArticleParams 检验文章参数
func checkArticleParams(params *models.ArticleSaveColumns) error {
	var valid validation.Validation
	valid.Required(params.Title, "title").Message("标题不能为空")
	valid.Required(params.Content, "content").Message("内容不能为空")
	valid.Required(params.Sort, "sort").Message("排序不能为空")
	valid.Required(params.CategoryId, "category_id").Message("分类ID不能为空")
	valid.MaxSize(params.Title, 80, "title").Message("标题长度不能超过80个字")
	valid.Min(params.Sort, 0, "sort").Message("排序不能是小于0的数字")
	hasError, msg := util.GetValidationMessage(valid)
	if hasError {
		return errors.New(msg)
	}
	return nil
}
