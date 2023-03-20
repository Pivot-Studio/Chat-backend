package client

import (
	"chat/common"
	"sync"
)

type Room struct {
	Members map[string]*Client
	MemLock sync.RWMutex
	Message chan []byte // 群组待发送消息
}

func NewRoom() *Room {
	room := &Room{
		Members: make(map[string]*Client),
		MemLock: sync.RWMutex{},
		Message: make(chan []byte, common.ClientSendDataBufferSize),
	}
	// 为Room创建一个协程, 用于把消息发送到群组内的所有用户
	go func() {
		for {
			select {
			case data := <-room.Message:
				room.SendDataToRoom(data)
			}
		}
	}()
	return room
}

func (r *Room) SendDataToRoom(data []byte) {
	r.MemLock.RLock()
	// todo: 根据消息类型, 进行不同的处理
	// todo: 避免群组内的用户给自己发消息
	for _, cli := range r.Members {
		cli.UnsentData <- data
	}
	r.MemLock.RUnlock()
}
