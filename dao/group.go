package dao

import (
	"chat/model"
	"gorm.io/gorm"
)

func (db *DBService) GetMemberGroupID(UserID uint) (GroupID []uint, err error) {
	// todo: 用redis缓存
	err = db.mysql.Model(&model.GroupMember{}).
		Select("group_id").
		Where("user_id = ?", UserID).
		Find(&GroupID).Error
	return
}

func (db *DBService) CreateGroup(group *model.Group) error {
	return db.mysql.Table("groups").Create(group).Error
}

func (db DBService) GetGroupByID(id uint) (group *model.Group, err error) {
	err = db.mysql.Table("groups").
		Where("id = ?", id).
		First(group).Error
	return
}

func (db DBService) AddNewUserToGroup(user *model.User, group *model.Group, role int) error {
	groupMember := &model.GroupMember{
		GroupID: group.ID,
		UserID:  user.ID,
		Role:    role,
	}
	err := db.mysql.
		Table("GroupMembers").
		Create(groupMember).Error
	return err
}

func (db DBService) GetGroupUsers(groupID uint) (members *[]model.GroupMember, err error) {
	err = db.mysql.
		Table("GroupMembers").
		Where("group_id = ?", groupID).
		Find(members).Error
	return
}

func (db DBService) IncrGroupUserNum(groupID uint) error {
	err := db.mysql.Table("groups").
		Where("group_id = ?", groupID).
		Update("mem_num", gorm.Expr("mem_num + 1")).Error
	return err
}
