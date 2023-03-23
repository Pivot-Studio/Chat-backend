package client

import (
	"chat/common"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"time"
)

type GroupMessageInput struct {
	SenderID   uint
	ReceiverID uint
	Content    string
	Type       int
	ReplyID    uint
}

type GroupMessageOutput struct {
	SenderID   uint
	ReceiverID uint
	Content    string
	Type       int
	ReplyID    uint
	SendTime   int64
}

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
	msg := &GroupMessageInput{}
	err := json.Unmarshal(data, msg)
	if err != nil {
		logrus.Errorf("Server:HandleGroupMessage:Unmarshal: %v", err)
		return
	}

	switch msg.Type {
	case common.TypeHeartBeat:
		c.HandleHeartBeat()
	case common.TypeGroupMessage:
		//todo: HandleGroupMessage这里的参数应该是一个结构体, 包含消息类型, 发送者信息, 消息内容等
		c.HandleGroupMessage(msg)
	}
}

func (c *Client) HandleGroupMessage(msg *GroupMessageInput) {
	// 根据消息查出对应的房间
	roomID := msg.ReceiverID
	CM.RoomLock.RLock()
	room := CM.Rooms[roomID]
	CM.RoomLock.RUnlock()

	// 重新封装消息, 带上发送者的信息, 以便群组内的其他用户知道是谁发的, 并且避免群组内的用户给自己发消息
	OutPut := &GroupMessageOutput{
		SenderID:   msg.SenderID,
		ReceiverID: msg.ReceiverID,
		Content:    msg.Content,
		Type:       msg.Type,
		ReplyID:    msg.ReplyID,
		SendTime:   time.Now().Unix(),
	}
	// 将消息发送到房间
	room.Message <- OutPut
}

func (c *Client) HandleHeartBeat() {
	c.HeartBeatTime = time.Now().Unix()
}
