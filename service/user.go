package service

import (
	"chat/dao"
	"chat/model"
	"github.com/sirupsen/logrus"
)

func Register(user *model.User) error {
	err := dao.DB.CreateUser(user)
	if err != nil {
		logrus.Errorf("[service.user.Register] %v", err)
		return err
	}
	return nil
}
