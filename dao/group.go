package dao

import "chat/model"

func (db *DBService) GetMemberGroupID(UserID uint) (GroupID []uint, err error) {
	// todo: 用redis缓存
	err = db.mysql.Model(&model.GroupMember{}).
		Select("group_id").
		Where("user_id = ?", UserID).
		Find(&GroupID).Error
	return
}

func (db *DBService) CreateGroup(group *model.Group) error {
	return db.mysql.Create(group).Error
}
