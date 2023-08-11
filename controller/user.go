package controller

import (
	"chat/common"
	"chat/model"
	"chat/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Register(c *gin.Context) {
	param := &model.RegisterParam{}
	err := c.ShouldBind(param)
	if err != nil {
		msg := "参数解析失败"
		logrus.Errorf("[controller.user.Register] %v", err)
		Response(c, http.StatusBadRequest, msg, nil)
		return
	}
	logrus.Infof("[controller.user.Register] %+v:注册请求", param)

	//todo:用户名，密码，邮箱，邀请码格式要求

	//密码加密
	passwordHash, err := common.EncodePassword(param.Password)
	param.Password = passwordHash
	if err != nil {
		logrus.Errorf("[controller.user.Register] %v", err)
		Response(c, http.StatusBadRequest, "密码Encode失败,注册失败", nil)
		return
	}

	err = service.Register(param)
	if err != nil {
		logrus.Errorf("[controller.user.Register] %v", err)
		Response(c, http.StatusBadRequest, "注册失败", nil)
		return
	}

	logrus.Infof("[controller.user.Register] %+v:注册成功", param)
	Response(c, http.StatusOK, "注册成功", nil)
}
