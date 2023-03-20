package model

import "gorm.io/gorm"

// Message 消息
type Message struct {
	gorm.Model
	SenderID   uint
	ReceiverID uint `gorm:"receiver_id, not null, index"`
	Content    string
	Type       int
	ReplyTo    uint
}
