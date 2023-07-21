package util

import (
	"gorm.io/gorm"
	"myblog/settings"
	"strconv"
)

// getOffset 获取偏移量
func getOffset(page, pageSize int) int {
	result := 0

	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}

// Pagination 分页程序
func Pagination(db *gorm.DB, page, pageSize int) *gorm.DB {
	return db.Limit(pageSize).Offset(getOffset(page, pageSize))
}

// Page 分页数据
func Page(db *gorm.DB, page, pageSize int) (pageResult map[string]interface{}) {
	var count int64
	db.Count(&count)
	pageResult = map[string]interface{}{
		"count":   count,
		"current": page,
	}
	if int(count)%pageSize == 0 {
		pageResult["total"] = int(count) / pageSize
	} else {
		pageResult["total"] = int(count)/pageSize + 1
	}
	return
}

// handlePageParams 处理分页参数
func handlePageParams(params map[string]string) (page int, pageSize int) {
	page = 1
	pageSize = settings.ServerConf.PageSize

	if _, ok := params["page"]; ok {
		page, _ = strconv.Atoi(params["page"])
	}

	if _, ok := params["page_size"]; ok {
		pageSize, _ = strconv.Atoi(params["page_size"])
	}
	return
}
