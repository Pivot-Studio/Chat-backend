package controller

import (
	"chat/common"
	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, code int, msg string, data map[string]interface{}) {
	if msg == "" {
		msg = common.GetErrorMessage(uint32(code), "")
	}

	if data == nil {
		data = make(map[string]interface{})
	}
	
	c.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
