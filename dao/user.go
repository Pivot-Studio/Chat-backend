package dao

import "chat/model"

func (db *DBService) CreateUser(user *model.User) error {
	return db.mysql.Create(user).Error
}
