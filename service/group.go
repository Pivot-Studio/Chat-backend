package service

import (
	"chat/dao"
	"chat/model"
	"github.com/sirupsen/logrus"
)

func CreateGroup(param *model.CreateGroupParam) error {
	user, err := dao.DB.GetUserByName(param.UserName)
	if err != nil {
		logrus.Errorf("[service.client.CreateGroup] %v", err)
		return err
	}

	group := &model.Group{
		OwnerID:      user.ID,
		Name:         param.GroupName,
		Introduction: param.Introduction,
	}
	//创建新群
	err = dao.DB.CreateGroup(group)
	if err != nil {
		logrus.Errorf("[service.group.CreateGroup] %v", err)
		return err
	}
	//写入G-U表
	err = dao.DB.AddNewUserToGroup(user, group, model.OWNER)
	if err != nil {
		logrus.Errorf("[service.group.CreateGroup] %v", err)
		return err
	}

	return nil
}

func JoinGroup(param *model.JoinGroupParam) error {
	user, err := dao.DB.GetUserByName(param.UserName)
	if err != nil {
		logrus.Errorf("[service.client.JoinGroup] %v", err)
		return err
	}

	group, err := dao.DB.GetGroupByID(param.GroupID)
	if err != nil {
		logrus.Errorf("[service.group.JoinGroup] %v", err)
		return err
	}

	err = dao.DB.AddNewUserToGroup(user, group, model.SPEAKER)
	if err != nil {
		logrus.Errorf("[service.group.JoinGroup] %v", err)
		return err
	}

	return nil
}
