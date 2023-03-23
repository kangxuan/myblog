package dao

import (
	"myblog/models"
	"time"
)

// GetCategoryList 获取分类列表
func GetCategoryList(maps map[string]interface{}) (categoryList []*models.Category, err error) {
	var args []interface{}
	args = append(args, 0)

	sqlStr := "select category_id, category_name, parent_id, create_time from category where is_delete = ?"

	if _, ok := maps["category_name"]; ok && maps["category_name"] != "" {
		sqlStr += " and category_name like '%?%"
		args = append(args, maps["category_name"])
	}

	if _, ok := maps["page_size"]; ok {
		sqlStr += " limit ?"
		args = append(args, maps["page_size"])
	}

	if _, ok := maps["page"]; ok {
		sqlStr += " offset ?"
		args = append(args, maps["page"])
	}

	err = models.Db.Select(&categoryList, sqlStr, args...)
	if err != nil {
		return nil, err
	}

	return categoryList, nil
}

// GetCategoryById 根据ID获取分类
func GetCategoryById(categoryId int) (category *models.Category, err error) {
	sqlStr := "select category_id, category_name, parent_id, create_time from category where is_delete = ? and category_id = ?"

	err = models.Db.Get(&category, sqlStr, 0, categoryId)
	if err != nil {
		return nil, err
	}

	return category, nil
}

// InsertCategory 插入分类
func InsertCategory(category models.Category) (int64, error) {
	currentTime := time.Now().Unix()
	sqlStr := "insert into category(category_name, category_type, parent_id, create_time, update_time) values (:category_name, :category_type, :parent_id, :create_time, :update_time)"
	result, err := models.Db.NamedExec(sqlStr, map[string]interface{}{
		"category_name": category.CategoryName,
		"category_type": category.CategoryType,
		"parent_id":     category.ParentId,
		"create_time":   currentTime,
		"update_time":   currentTime,
	})
	if err != nil {
		return 0, err
	}
	lastCategoryId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastCategoryId, nil
}
