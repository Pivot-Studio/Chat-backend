package controller

import (
	"chat/model"
	"chat/service/client"
	"chat/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//wjc
type registerParam struct {
	Username        string `form:"username" binding:"required"`
	Password        string `form:"password" binding:"required"`
	ConfirmPassword string `form:"confirmpassword" binding:"required"`
	Email           string `form:"email" binding:"required"`
	InvitationCode  string `form:"invitationcode"`
}

func Register(c *gin.Context) {
	p := &registerParam{}
	err := c.ShouldBind(p)

	if err != nil {
		msg := "参数解析失败,注册失败"
		logrus.Errorf("[controller.user.Register] %+v", err)
		Response(c, http.StatusBadRequest, msg, nil)
		return
	}

	if p.Password != p.ConfirmPassword {
		msg := "确认密码与先前输入不一致，注册失败"
		logrus.Errorf("[controller.user.Register] ConfirmPassword is not identical")
		Response(c, http.StatusBadRequest, msg, nil)
		return
	}

	//todo:用户名，密码，邮箱，邀请码格式要求

	//密码加密
	passwordHash, err := util.EncodePassword(p.Password)

	if err != nil {
		msg := "密码Encode失败,注册失败"
		logrus.Errorf("[controller.user.Register] %+v", err)
		Response(c, http.StatusBadRequest, msg, nil)
		return
	}

	//写入Mysql数据库
	err = client.Register(&model.User{
		Username:       p.Username,
		Email:          p.Email,
		Avatar:         "",
		InvitationCode: p.InvitationCode,
		Password:       passwordHash,
	})

	if err != nil {
		msg := "写入mysql失败,注册失败"
		logrus.Errorf("[controller.user.Register] %+v", err)
		Response(c, http.StatusBadRequest, msg, nil)
		return
	}

	logrus.Infof("[controller.user.Register] 注册成功")
	Response(c, http.StatusOK, "", nil)
}
