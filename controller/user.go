package controller

import (
	"chat/common"
	"chat/model"
	"chat/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// wjc
type registerParam struct {
	Username       string `form:"username" binding:"required"`
	Password       string `form:"password" binding:"required"`
	Email          string `form:"email" binding:"required"`
	InvitationCode string `form:"invitation_code"`
}

func Register(c *gin.Context) {
	param := &registerParam{}
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

	if err != nil {
		logrus.Errorf("[controller.user.Register] %v", err)
		Response(c, http.StatusBadRequest, "密码Encode失败,注册失败", nil)
		return
	}

	//写入Mysql数据库
	err = service.Register(&model.User{
		Username:       param.Username,
		Email:          param.Email,
		Avatar:         "",
		InvitationCode: param.InvitationCode,
		Password:       passwordHash,
	})

	if err != nil {
		logrus.Errorf("[controller.user.Register] %v", err)
		Response(c, http.StatusBadRequest, "注册失败", nil)
		return
	}

	logrus.Infof("[controller.user.Register] %+v:注册成功", param)
	Response(c, http.StatusOK, "注册成功", nil)
}
