package model

import "gorm.io/gorm"

// User 用户表
type User struct {
	gorm.Model
	UserName       string
	Password       string
	Email          string
	InvitationCode string
}
