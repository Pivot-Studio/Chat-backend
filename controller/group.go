package controller

import (
	"chat/model"
	"chat/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateGroup(c *gin.Context) {
	p := &model.CreateGroupParam{}
	err := c.ShouldBindJSON(p)
	if err != nil || p.GroupName == "" {
		msg := "参数不合法,创建群组失败"
		logrus.Errorf("[controller.group.CreatGroup] %v", err)
		Response(c, http.StatusBadRequest, msg, nil)
	}

	err = service.CreateGroup(p)
	if err != nil {
		msg := "创建群组失败"
		logrus.Errorf("[controller.group.CreateGroup] %v", err)
		Response(c, http.StatusBadRequest, msg, nil)
	}

	logrus.Infof("[controller.group.CreatGroup] 创建群组成功")
	Response(c, http.StatusOK, "", nil)
}

func JoinGroup(c *gin.Context) {
	p := &model.JoinGroupParam{}
	err := c.ShouldBindJSON(p)
	if err != nil {
		msg := "参数不合法,加入群组失败"
		logrus.Errorf("[controller.group.JoinGroup] %v", err)
		Response(c, http.StatusBadRequest, msg, nil)
	}

	err = service.JoinGroup(p)
	if err != nil {
		msg := "加入群组失败"
		logrus.Errorf("[controller.group.JoinGroup] %v", err)
		Response(c, http.StatusBadRequest, msg, nil)
	}

	logrus.Infof("[controller.group.JoinGroup] 加入群组成功")
	Response(c, http.StatusOK, "", nil)
}
