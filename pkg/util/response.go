package util

import (
	"github.com/gin-gonic/gin"
	"myblog/pkg/e"
)

type Gin struct {
	C *gin.Context
}

func (c *Gin) Response(httpCode int, errCode int, data interface{}) {
	c.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})
}
