package dao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"myblog/models"
	"myblog/pkg/util"
	"time"
)

type CategoryColumns struct {
	CategoryId    int    `json:"category_id"`
	CategoryType  int    `json:"category_type"`
	CategoryName  string `json:"category_name"`
	ParentId      int    `json:"parent_id"`
	CreateTime    int    `json:"create_time"`
	UpdateTime    int    `json:"update_time"`
	CreateTimeStr string `json:"create_time_str"`
	UpdateTimeStr string `json:"update_time_str"`
}

// GetCategoryList 获取分类列表
func GetCategoryList(params map[string]string) (categoryList []*CategoryColumns, err error) {
	query := getCategoryQuery(params)
	util.Pagination(query, params).Find(&categoryList)
	for k, cat := range categoryList {
		categoryList[k].CreateTimeStr = util.TimeToString(cat.CreateTime)
		categoryList[k].UpdateTimeStr = util.TimeToString(cat.UpdateTime)
	}

	return categoryList, nil
}

// GetCategoryPage 获取分类分页
func GetCategoryPage(params map[string]string) (page map[string]interface{}) {
	fmt.Println(params)
	query := getCategoryQuery(params)
	page = util.Page(query, params)
	return
}

// getCategoryQuery 分类通用的查询器
func getCategoryQuery(params map[string]string) *gorm.DB {
	whereMaps := map[string]interface{}{}
	whereMaps["is_delete"] = 0
	if _, ok := params["category_name"]; ok && params["category_name"] != "" {
		whereMaps["category_name"] = params["category_name"]
	}

	return models.Db.Table("category").Where(whereMaps).Order("category_id desc")
}

// GetCategoryById 根据ID获取分类
func GetCategoryById(categoryId int) (category *CategoryColumns, err error) {
	result := models.Db.Table("category").Limit(1).Where("category_id = ?", categoryId).Find(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	category.CreateTimeStr = util.TimeToString(category.CreateTime)
	category.UpdateTimeStr = util.TimeToString(category.UpdateTime)
	return
}

// InsertCategory 插入分类
func InsertCategory(categoryParams models.Category) (int, error) {
	// 判断分类名称是否已经存在
	existedNum := models.Db.Table("category").Where("category_name = ?", categoryParams.CategoryName).Find(&models.Category{}).RowsAffected
	if existedNum > 0 {
		err := errors.New("分类名称已存在")
		return 0, err
	}
	// 新增数据
	currentTime := time.Now().Unix()
	category := models.Category{
		CategoryName: categoryParams.CategoryName,
		CategoryType: categoryParams.CategoryType,
		ParentId:     categoryParams.ParentId,
		CreateTime:   int(currentTime),
		UpdateTime:   int(currentTime),
	}
	result := models.Db.Table("category").Create(&category)
	if result.Error != nil {
		return 0, result.Error
	}
	return category.CategoryId, nil
}
