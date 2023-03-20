package client

import (
	"chat/common"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"time"
)

func (c *Client) ClientSendBiData(data []byte) {
	_ = c.Socket.SetWriteDeadline(time.Now().Add(common.ClientWriteDataTimeout))

	err := c.Socket.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		logrus.Errorf("Server:ClientSendBiData:WriteMessage: %v", err)
		// 断开连接, 至于连接断开后的处理, 交给ClientManager异步清扫
		c.DeleteClient()
	}
}

func (c *Client) HandleClientMessage(data []byte) {
	// todo: 根据消息类型, 调用不同的处理函数
	Type := uint(1)

	switch Type {
	case common.TypeHeartBeat:
		c.HandleHeartBeat()
	case common.TypeGroupMessage:
		//todo: HandleGroupMessage这里的参数应该是一个结构体, 包含消息类型, 发送者信息, 消息内容等
		c.HandleGroupMessage(data)
	}
}

func (c *Client) HandleGroupMessage(data []byte) {
	// todo: 根据消息查出对应的房间, 并将消息发送到房间的消息队列中
	roomID := uint(1)

	CM.RoomLock.RLock()
	room := CM.Rooms[roomID]
	CM.RoomLock.RUnlock()

	// todo: 重新封装消息, 注意要带上发送者的信息, 以便群组内的其他用户知道是谁发的, 并且避免群组内的用户给自己发消息

	OutPut := []byte("hello world")
	room.Message <- OutPut
}

func (c *Client) HandleHeartBeat() {
	c.HeartBeatTime = time.Now().Unix()
}
