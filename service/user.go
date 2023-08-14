package service

import (
	"chat/dao"
	"chat/model"
	"github.com/sirupsen/logrus"
)

func Register(param *model.RegisterParam) error {
	user := &model.User{
		Username:       param.Username,
		Email:          param.Email,
		Avatar:         "",
		InvitationCode: param.InvitationCode,
		Password:       param.Password,
	}
	err := dao.DB.CreateUser(user)
	if err != nil {
		logrus.Errorf("[service.user.Register] %v", err)
		return err
	}
	return nil
}
