package main

import (
	"chat/controller"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	webMode := viper.Get("app.web_mode")
	webPort := viper.Get("app.web_port")
	if webMode == "gin" {
		// gin 作为web服务
		r := gin.Default()
		controller.Router(r)
		controller.Engine.Run(fmt.Sprintf(":%v", webPort))
	} else {
		//todo: rpc

	}

}
