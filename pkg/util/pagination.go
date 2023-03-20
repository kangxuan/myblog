package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"myblog/settings"
)

// GetPage 获取页码
func GetPage(c *gin.Context) int {
	result := 0

	if arg := c.Query("page"); arg != "" {
		page, err := com.StrTo(c.Query("page")).Int()
		if err != nil {
			panic(fmt.Errorf("StrTo Err: %v", err))
		}
		if page > 0 {
			result = (page - 1) * settings.ServerConf.PageSize
		}
	}

	return result
}
