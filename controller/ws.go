package controller

import (
	"chat/common"
	"chat/service/client"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)
import "github.com/gorilla/websocket"

var WsUpgrader = websocket.Upgrader{
	HandshakeTimeout: common.WebsocketHandshakeTimeout,
	ReadBufferSize:   common.WebSocketReadBufferSize,
	WriteBufferSize:  common.WebSocketWriteBufferSize,
	CheckOrigin: func(r *http.Request) bool {
		// todo:用于鉴别跨站请求, 检查Origin字段
		return true
	},
}

func WsHandler(c *gin.Context) {
	// todo: auth token, 获取AppID和UserID
	// todo: 判断在线等操作
	var AppID uint = 1
	var UserID uint = 1

	ClientIP := c.ClientIP()
	con, err := WsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logrus.Errorf("controller:WsHandler %v, ClientIP:%v", err, ClientIP)
		return
	}
	logrus.Infof("controller:WsHandler, Client IP:%v", ClientIP)

	// 登录成功, 创建连接
	go client.CM.ClientLogin(ClientIP, AppID, UserID, con)

	// todo:返回前端
}
