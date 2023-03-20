package common

import "time"

const (
	WebsocketHandshakeTimeout = 10 * time.Second
	WebSocketReadBufferSize   = 1024
	WebSocketWriteBufferSize  = 1024
)

const (
	ClientReadDataTimeout    = 60 * time.Second
	ClientWriteDataTimeout   = 60 * time.Second
	ClientHeartBeatTimeout   = 10 * time.Second
	ClientSendDataBufferSize = 1024
	ClientLoginBufferSize    = 100
	ClientLogoutBufferSize   = 100
)

const (
	ClientMessageSweepInterval = 10 * time.Second
)

const (
	TypeHeartBeat    = 0
	TypeGroupMessage = 1
)
