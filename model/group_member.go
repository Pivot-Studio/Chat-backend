package model

import "gorm.io/gorm"

// GroupMember 群组成员
type GroupMember struct {
	gorm.Model
	GroupID uint `gorm:"group_id, not null, index"`
	UserID  uint `gorm:"user_id, not null, index"`
	Role    int  // 用户在当前群组的role
	Status  int  // 状态, 上线/离线
}
