package client

import (
	"chat/common"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"time"
)

type Client struct {
	Addr          string
	AppID         uint
	UserID        uint
	Socket        *websocket.Conn
	UnsentData    chan []byte // 待发送数据
	HeartBeatTime int64
	LoginTime     int64
}

func (c *Client) GetKey() string {
	// Key = UserID_AppID
	return fmt.Sprintf("%v_%v", c.UserID, c.AppID)
}

func NewClient(ClientIP string, AppID uint, UserID uint, con *websocket.Conn) *Client {
	currentTime := time.Now().Unix()
	cli := &Client{
		Addr:          ClientIP,
		AppID:         AppID,
		UserID:        UserID,
		Socket:        con,
		UnsentData:    make(chan []byte, common.ClientSendDataBufferSize),
		HeartBeatTime: currentTime,
		LoginTime:     currentTime,
	}

	// 为Client创建一个协程, 用于发送数据
	go func() {
		for {
			select {
			case data := <-cli.UnsentData:
				cli.ClientSendBiData(data)
			}
		}
	}()

	// 监听客户端发送的数据
	go func() {
		for {
			_ = cli.Socket.SetReadDeadline(time.Now().Add(common.ClientReadDataTimeout))
			_, data, err := cli.Socket.ReadMessage()
			if err != nil {
				logrus.Errorf("Server:ClientReadData:ReadMessage: %v", err)
				// 断开连接, 至于连接断开后的处理, 交给ClientManager异步清扫
				cli.DeleteClient()
				return
			}
			cli.HandleClientMessage(data)
		}
	}()
	return cli
}

// DeleteClient 断开连接, 并令Socket=nil, 但不删除Client, 交给ClientManager异步清扫
func (c *Client) DeleteClient() {
	if c.Socket != nil {
		_ = c.Socket.Close()
		c.Socket = nil
	}
}

// IsOnline 判断客户端是否在线, 通过判断Socket是否为nil和心跳时间来判断
func (c *Client) IsOnline() bool {
	// 断开连接或心跳超时, 则认为不在线
	// todo: 判断心跳超时, 先不做
	if c.Socket == nil {
		return false
	}
	return true
}
