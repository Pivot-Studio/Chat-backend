package client

import (
	"chat/common"
	"chat/dao"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var CM *ClientManager

type ClientManager struct {
	Clients    map[string]*Client
	ClientLock sync.RWMutex

	Rooms    map[uint]*Room
	RoomLock sync.RWMutex

	Login  chan *Client
	Logout chan *Client
}

func init() {
	CM = &ClientManager{
		Clients:    make(map[string]*Client),
		ClientLock: sync.RWMutex{},
		Rooms:      make(map[uint]*Room),
		RoomLock:   sync.RWMutex{},
		Login:      make(chan *Client, 100),
		Logout:     make(chan *Client, 100),
	}

	// 定时任务, 清理已经断开连接的客户端
	go func() {
		for {
			CM.SweepLogoutClient()
			time.Sleep(common.ClientMessageSweepInterval)
		}
	}()

	// 监听任务, 处理客户端的登录与退出
	// 由于清扫任务的存在, 会导致一段时间内, 无法进行登录与退出的处理, 会被阻塞
	// 所以采用channel的方式, 异步处理
	go func() {
		for {
			select {
			case cli := <-CM.Login:
				LoginEvent(cli)
			case cli := <-CM.Logout:
				LogoutEvent(cli)
			}
		}
	}()
}

func (cm *ClientManager) ClientLogin(ClientIP string, AppID uint, UserID uint, con *websocket.Conn) {
	cli := NewClient(ClientIP, AppID, UserID, con)
	CM.Login <- cli
}

func (cm *ClientManager) ClientLogout(client *Client) {
	CM.Logout <- client
}

// SweepLogoutClient 定时清理已经断开连接的客户端
func (cm *ClientManager) SweepLogoutClient() {
	cm.ClientLock.Lock()
	for _, client := range cm.Clients {
		if client.IsOnline() {
			LogoutEvent(client)
			// todo: 通知其他客户端, 该用户已经下线

		}
	}
	cm.ClientLock.Unlock()
}

// LoginEvent 登录事件
func LoginEvent(client *Client) {
	// 这里允许Clients与Rooms的不同步
	// 一般而言, Clients先于Rooms创建, 后于Rooms删除

	// 加入连接列表
	CM.ClientLock.Lock()
	CM.Clients[client.GetKey()] = client
	CM.ClientLock.Unlock()

	// 查库得到用户所在的roomId
	rooms, err := dao.DB.GetMemberGroupID(client.UserID)
	if err != nil {
		return
	}

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

// LogoutEvent 退出事件
func LogoutEvent(client *Client) {
	// 退出rooms
	roomIDs, err := dao.DB.GetMemberGroupID(client.UserID)
	if err != nil {
		logrus.Errorf("LogoutEvent:GetMemberGroupID ClientID:%v_%v, error: %v",
			client.UserID, client.AppID, err)
		return
	}

	for _, roomID := range roomIDs {
		go func(ID uint) {
			CM.RoomLock.RLock()
			room := CM.Rooms[ID]
			defer CM.RoomLock.RUnlock()

			if room == nil {
				return
			}

			room.MemLock.Lock()
			delete(room.Members, client.GetKey())
			room.MemLock.Unlock()
			// todo: room为空时, 删除room
		}(roomID)
	}

	// 退出连接列表
	CM.ClientLock.Lock()
	delete(CM.Clients, client.GetKey())
	CM.ClientLock.Unlock()

	// 清除Client
	client.DeleteClient()
}

// IsOnline 判断用户是否在线, 注意与Client.IsOnline()的区别, 这里是判断是否存在Client
func (cm *ClientManager) IsOnline(AppID uint, UserID uint) bool {
	cm.ClientLock.RLock()
	defer cm.ClientLock.RUnlock()

	cli := cm.Clients[GetClientKey(UserID, AppID)]

	return cli != nil
}

// GetClient 获取Client
func (cm *ClientManager) GetClient(AppID uint, UserID uint) *Client {
	cm.ClientLock.RLock()
	defer cm.ClientLock.RUnlock()

	return cm.Clients[GetClientKey(UserID, AppID)]
}

// GetClientKey 获取Client的key
func GetClientKey(UserID, AppID uint) string {
	return fmt.Sprintf("%v_%v", UserID, AppID)
}
