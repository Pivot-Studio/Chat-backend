package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CreateGroupParam struct {
	GroupName    string `json:"group_name"`
	Introduction string `json:"introduction"`
}

func CreateGroup(c *gin.Context) {
	p := &CreateGroupParam{}
	err := c.ShouldBindJSON(p)
	if err != nil || p.GroupName == "" {
		msg := "参数不合法,创建群组失败"
		logrus.Errorf("[controller.group.CreatGroup] %+v", err)
		Response(c, http.StatusBadRequest, msg, nil)
	}

	logrus.Infof("[controller.group.CreatGroup] 注册成功")
	Response(c, http.StatusOK, "", nil)
}
