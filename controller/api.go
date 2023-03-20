package controller

import "github.com/gin-gonic/gin"

var Engine *gin.Engine

func Router(r *gin.Engine) {
	r.GET("/login")
}
