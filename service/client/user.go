package client

import (
	"chat/dao"
	"chat/model"
)

func Register(user *model.User) error {
	return dao.RS.CreateUser(user)
}
