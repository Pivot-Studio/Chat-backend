package common

import "time"

// websocket
const (
	WebsocketHandshakeTimeout = 10 * time.Second
	WebSocketReadBufferSize   = 1024
	WebSocketWriteBufferSize  = 1024
)

// client设置
const (
	ClientReadDataTimeout    = 60 * time.Second
	ClientWriteDataTimeout   = 60 * time.Second
	ClientHeartBeatTimeout   = 10 * time.Second
	ClientSendDataBufferSize = 1024
	ClientLoginBufferSize    = 100
	ClientLogoutBufferSize   = 100
)

// client管理相关
const (
	ClientMessageSweepInterval = 10 * time.Second
)

// 消息类型
const (
	TypeHeartBeat    = 0
	TypeGroupMessage = 1
)

// 群组成员
const (
	GroupMemberRoleOwner = 1
	GroupMemberRoleAdmin = 2
	GroupMemberRoleUser  = 3
)

const (
	AppPlatformWeb     = 1
	AppPlatformIOS     = 2
	AppPlatformAndroid = 3
	AppPlatformMac     = 4
)
