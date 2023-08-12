package service

import (
	"chat/dao"
	"chat/model"
	"github.com/sirupsen/logrus"
	"sync"
)

type Group_ struct {
	Group   *model.Group
	Members *[]model.GroupMember
	sync.RWMutex
}

type GroupOperator struct {
	GroupsMap sync.Map
	lock      sync.Mutex
}

var GroupOp GroupOperator

func (gpo GroupOperator) StoreGroup(groupID uint, group *Group_) {
	GroupOp.GroupsMap.Store(groupID, group)
}

func (gpo GroupOperator) GetGroup(groupID uint) (*Group_, error) {
	gpo.lock.Lock()
	defer gpo.lock.Unlock()

	v, ok := gpo.GroupsMap.Load(groupID)
	if !ok {
		g, err := dao.DB.GetGroupByID(groupID)
		if err != nil {
			logrus.Errorf("[service.group_operator.GetGroup] %v", err)
			return nil, err
		}
		v = &Group_{
			Group: g,
		}
		gpo.StoreGroup(groupID, v.(*Group_))
	}

	g := v.(*Group_)
	g.Lock()
	if g.Group.MemNum != uint(len(*g.Members)) {
		var err error
		g.Members, err = dao.DB.GetGroupUsers(groupID)
		if err != nil {
			logrus.Errorf("[service.group_operator.GetGroup] %v", err)
			g.Unlock()
			return nil, err
		}
	}
	g.Unlock()
	return g, nil
}
