package dao

import "github.com/gin-gonic/examples/onboarding/po"

// 数据层操作
type userDaoI interface {
	// 保存数据
	Add(user po.User) po.User
	// 获取全部用户
	QueryAll() []po.User
}
