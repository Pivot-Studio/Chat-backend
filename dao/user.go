package dao

import "chat/model"

func (db *DBService) CreateUser(user *model.User) error {
	return db.mysql.Create(user).Error
}

/*
func CreateUser(user *model.User) error {
	err := db.Table("users").Create(user).Error
	return err
}
*/

func (db *DBService) GetUserByName(username string) (user *model.User, err error) {
	err = db.mysql.Table("users").Where("username = ?", username).First(user).Error
	return
}
