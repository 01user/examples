package dao

import "github.com/gin-gonic/examples/onboarding/po"

// 数据层操作
type relationshipDaoI interface {
	// 保存数据
	Add(rel po.Relationship) po.Relationship
	// 保存数据，处理匹配关系
	AddAll(rels []po.Relationship) []po.Relationship
	// 获取全部
	QueryAll(userId int) []po.Relationship
	// 获取单个
	GetOne(userId int, otherUserId int) (po.Relationship, bool)
}
