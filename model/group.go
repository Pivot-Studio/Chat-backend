package model

import "gorm.io/gorm"

type CreateGroupParam struct {
	UserName     string `json:"user_name"`
	GroupName    string `json:"group_name"`
	Introduction string `json:"introduction"`
}

// Group 群组
type Group struct {
	gorm.Model
	OwnerID      uint `gorm:"owner_id"`
	Name         string
	Avatar       string
	Introduction string
	MemNum       uint
}

type JoinGroupParam struct {
	UserName string `json:"user_name"`
	GroupID  uint   `json:"group_id" binding:"required"`
}
