package model

import "gorm.io/gorm"

// User 用户表
type User struct {
	gorm.Model
	Username       string `gorm:"index"`
	Password       string
	Avatar         string
	Email          string `gorm:"uniqueIndex"`
	InvitationCode string
}
