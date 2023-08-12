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
		MemNum:       1,
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

	g, err := GroupOp.GetGroup(param.GroupID)
	if err != nil {
		logrus.Errorf("[service.group.JoinGroup] %v", err)
		return err
	}

	g.Lock()
	defer g.Unlock()

	err = dao.DB.AddNewUserToGroup(user, g.Group, model.SPEAKER)
	if err != nil {
		logrus.Errorf("[service.group.JoinGroup] %v", err)
		return err
	}
	err = dao.DB.IncrGroupUserNum(g.Group.ID)
	if err != nil {
		logrus.Errorf("[service.group.JoinGroup] %v", err)
		return err
	}

	groupMember := model.GroupMember{
		GroupID: g.Group.ID,
		UserID:  user.ID,
		Role:    model.SPEAKER,
	}
	*g.Members = append(*g.Members, groupMember)
	g.Group.MemNum += 1

	return nil
}
