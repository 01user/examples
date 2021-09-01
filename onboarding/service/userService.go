package service

import (
	"github.com/gin-gonic/examples/onboarding/dao"
	"github.com/gin-gonic/examples/onboarding/po"
)

type UserService struct {
	Id int
}

// 查询关系
func (userService UserService) QueryAll() []po.User {
	userDao := new(dao.UserDao)
	users := userDao.QueryAll()
	return users
}

// 添加用户
func (userService UserService) AddUser(user po.User) po.User {
	userDao := new(dao.UserDao)
	add := userDao.Add(user)
	return add
}
