package dao

import "chat/model"

func (rs *RdbService) CreateUser(user *model.User) error {
	return rs.tx.Create(&user).Error
}
