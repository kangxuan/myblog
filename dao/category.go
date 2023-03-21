package dao

import (
	"myblog/models"
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
