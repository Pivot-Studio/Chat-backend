package client

import (
	"chat/common"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

var CM *ClientManager

type ClientManager struct {
	Clients    map[string]*Client
	ClientLock sync.RWMutex

	Rooms    map[uint]*Room
	RoomLock sync.RWMutex
}

func init() {
	CM = &ClientManager{
		Clients:    make(map[string]*Client),
		ClientLock: sync.RWMutex{},
		Rooms:      make(map[uint]*Room),
		RoomLock:   sync.RWMutex{},
	}

	// 定时任务, 清理已经断开连接的客户端
	go func() {
		for {
			CM.SweepLogoutClient()
			time.Sleep(common.ClientMessageSweepInterval)
		}
	}()
}

func (cm *ClientManager) ClientLogin(ClientIP string, AppID uint, UserID uint, con *websocket.Conn) {
	cli := NewClient(ClientIP, AppID, UserID, con)
	LoginEvent(cli)
}

func (cm *ClientManager) ClientLogout(client *Client) {
	LogoutEvent(client)
}

// SweepLogoutClient 定时清理已经断开连接的客户端
func (cm *ClientManager) SweepLogoutClient() {
	cm.ClientLock.Lock()
	for _, client := range cm.Clients {
		if client.IsOnline() {
			LoginEvent(client)
		}
	}
	cm.ClientLock.Unlock()
}

func LoginEvent(client *Client) {
	// 这里允许Clients与Rooms的不同步
	// 一般而言, Clients先于Rooms创建, 后于Rooms删除

	// 加入连接列表
	CM.ClientLock.Lock()
	CM.Clients[client.GetKey()] = client
	CM.ClientLock.Unlock()

	// todo: 查库, 有哪些群
	var rooms []uint

	// 加入room
	for _, roomID := range rooms {
		go func(ID uint) {
			// 查看room否存在
			CM.RoomLock.RLock()
			room := CM.Rooms[ID]
			CM.RoomLock.RUnlock()

			if room == nil {
				// 不存在则创建
				room = NewRoom()
				CM.RoomLock.Lock()
				CM.Rooms[ID] = room
				CM.RoomLock.Unlock()
			}

			// 加入room
			room.MemLock.Lock()
			room.Members[client.GetKey()] = client
			// todo: 这里发送可以同步的消息, 通知客户端同步群组信息?
			room.MemLock.Unlock()
		}(roomID)
	}
}

func LogoutEvent(client *Client) {
	// 退出room
	// todo: 查库, 有哪些群
	var roomIDs []uint

	for _, roomID := range roomIDs {
		go func(ID uint) {
			CM.RoomLock.RLock()
			room := CM.Rooms[ID]
			CM.RoomLock.RUnlock()

			if room == nil {
				return
			}

			room.MemLock.Lock()
			delete(room.Members, client.GetKey())
			room.MemLock.Unlock()
		}(roomID)
	}

	// 退出连接列表
	CM.ClientLock.Lock()
	delete(CM.Clients, client.GetKey())
	CM.ClientLock.Unlock()

	// 删除连接
	client.DeleteClient()
}
