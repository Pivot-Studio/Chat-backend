package model

import "gorm.io/gorm"

// Group 群组
type Group struct {
	gorm.Model
	OwnerID      uint `gorm:"group_id"`
	Name         string
	Introduction string
	MemNum       uint
}
