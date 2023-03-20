package controller

import "github.com/gin-gonic/gin"

func init() {
	r := gin.Default()
	Router(r)
}
