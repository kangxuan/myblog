package util

import (
	"github.com/gin-gonic/gin"
	"myblog/pkg/e"
)

type Gin struct {
	C *gin.Context
}

func (c *Gin) Response(httpCode int, errCode int, msg string, data interface{}) {
	if msg == "" {
		msg = e.GetMsg(errCode)
	}
	c.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  msg,
		"data": data,
	})
}

func Response(c *gin.Context, httpCode int, errCode int, msg string, data interface{}) {
	if msg == "" {
		msg = e.GetMsg(errCode)
	}
	c.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  msg,
		"data": data,
	})
}
