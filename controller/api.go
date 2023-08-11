package controller

import (
	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine

func Router(r *gin.Engine) {
	api := r.Group("/api")
	{
		user := api.Group("/user")
		{
			user.GET("/login", WsHandler)
			user.POST("/register", Register)
		}
		group := api.Group("/group")
		{
			group.POST("/create", CreateGroup)
			group.POST("/join", JoinGroup)
		}
	}
}
