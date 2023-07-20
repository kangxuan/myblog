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

// GetTagList 获取标签列表
func GetTagList(c *gin.Context) {
	appG := util.Gin{C: c}
	params := make(map[string]string)
	params["tag_name"] = c.DefaultQuery("tag_name", "")
	tagList, err := dao.GetTagList(params)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, "获取标签列表成功", tagList)
}

// CreateTag 创建标签
func CreateTag(c *gin.Context) {
	appG := util.Gin{C: c}

	var tagParams *models.Tag
	err := c.BindJSON(&tagParams)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, "参数解析错误", nil)
		return
	}

	err = checkTagParams(tagParams)
	if err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, err.Error(), nil)
		return
	}

	lastTagId, err := dao.CreateTag(tagParams)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, "添加标签成功", map[string]interface{}{
		"tag_id": lastTagId,
	})
}

// UpdateTag 更新标签
func UpdateTag(c *gin.Context) {
	appG := util.Gin{C: c}

	id, _ := strconv.Atoi(c.Param("id"))
	var tagParams *models.Tag
	err := c.BindJSON(&tagParams)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, "参数解析错误", nil)
		return
	}

	err = checkTagParams(tagParams)
	if err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, err.Error(), nil)
		return
	}

	err = dao.UpdateTag(id, tagParams)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, "修改标签成功", nil)
}

// DeleteTag 删除标签
func DeleteTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := dao.DeleteTag(id)
	if err != nil {
		util.Response(c, http.StatusOK, e.ERROR, err.Error(), nil)
		return
	}
	util.Response(c, http.StatusOK, e.SUCCESS, "删除标签成功", nil)
}

// checkTagParams 检查标签参数
func checkTagParams(tagParams *models.Tag) error {
	var valid validation.Validation
	valid.Required(tagParams.TagName, "tag_name").Message("标签名称不能为空")
	valid.MaxSize(tagParams.TagName, 50, "tag_name").Message("标签名称不能超过50个字")
	hasError, msg := util.GetValidationMessage(valid)
	if hasError {
		return errors.New(msg)
	}
	return nil
}
