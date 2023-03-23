package client

import (
	"chat/common"
	"encoding/json"
	"sync"
)

type Room struct {
	Members map[string]*Client
	MemLock sync.RWMutex
	Message chan *GroupMessageOutput // 群组待发送消息
}

func NewRoom() *Room {
	room := &Room{
		Members: make(map[string]*Client),
		MemLock: sync.RWMutex{},
		Message: make(chan *GroupMessageOutput, common.ClientSendDataBufferSize),
	}
	// 为Room创建一个协程, 用于把消息发送到群组内的所有用户
	go func() {
		for {
			select {
			case data := <-room.Message:
				room.SendDataToMembers(data)
			}
		}
	}()
	return room
}

func (r *Room) SendDataToMembers(msg *GroupMessageOutput) {
	data, _ := json.Marshal(msg)
	r.MemLock.RLock()
	// todo: 根据消息类型, 进行不同的处理
	for _, cli := range r.Members {
		// 避免群组内的用户给自己发消息
		if cli.UserID == msg.SenderID {
			continue
		}
		cli.UnsentData <- data
	}
	r.MemLock.RUnlock()
}
