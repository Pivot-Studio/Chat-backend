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

type RegisterParam struct {
	Username       string `form:"username" binding:"required"`
	Password       string `form:"password" binding:"required"`
	Email          string `form:"email" binding:"required,email"`
	InvitationCode string `form:"invitation_code"`
}
