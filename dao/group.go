package dao

import "chat/model"

func GetMemberGroupID(UserID uint) (GroupID []uint, err error) {
	// todo: 用redis缓存

	err = Mysql.Model(&model.GroupMember{}).
		Select("group_id").
		Where("user_id = ?", UserID).
		Find(&GroupID).Error
	return
}

func (rs *RdbService) CreateGroup(group *model.Group) error {
	return rs.tx.Create(&group).Error
}
