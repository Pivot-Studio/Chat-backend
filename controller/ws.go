package controller

import (
	"chat/common"
	"chat/service/client"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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
	Account := c.PostForm("account")
	Password := c.PostForm("password")
	AppIDRaw := c.PostForm("app_id")
	var UserID uint = 1

	if Account == "" || Password == "" || AppIDRaw == "" {
		Token, err := c.Cookie("token")
		if err != nil {
			Response(c, http.StatusUnauthorized, "", nil)
			return
		}

		// todo: 解析token, 获取AppID和UserID
		logrus.Infof("api:WsHandler, Token:%v", Token)
		return
	}

	AppID, err := strconv.ParseInt(AppIDRaw, 10, 64)
	if err != nil {
		Response(c, http.StatusBadRequest, "", nil)
		return
	}
	// todo: 判断AppID是否存在

	// 判断在线, 如果已经在线, 则断开旧连接, 重新登录
	if client.CM.IsOnline(uint(AppID), UserID) {
		// 断开并清理旧连接
		cli := client.CM.GetClient(uint(AppID), UserID)
		client.LogoutEvent(cli)
	}

	// todo: 在升级前, 生成token, 返回给客户端
	Response(c, http.StatusOK, "", nil)

	// 升级为websocket连接
	ClientIP := c.ClientIP()
	con, err := WsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logrus.Errorf("controller:WsHandler %v, ClientIP:%v", err, ClientIP)
		return
	}
	logrus.Infof("controller:WsHandler, Client IP:%v", ClientIP)

	// 创建连接
	client.CM.ClientLogin(ClientIP, uint(AppID), UserID, con)
}
