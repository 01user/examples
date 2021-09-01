package service

import (
	"github.com/gin-gonic/examples/onboarding/dao"
	"github.com/gin-gonic/examples/onboarding/po"
)

type RelationshipService struct {
	Id int
}

// 查询关系
func (relService RelationshipService) QueryAll(userId int) []po.Relationship {
	relDao := new(dao.RelationshipDao)
	relationships := relDao.QueryAll(userId)
	return relationships
}

// 添加关系
func (relService RelationshipService) AddRelationship(rel po.Relationship) po.Relationship {
	relationships := HandleRelationship(rel)
	relDao := new(dao.RelationshipDao)
	addAll := relDao.AddAll(relationships)
	for _, relationshipHandled := range addAll {
		if relationshipHandled.UserId == rel.UserId {
			return relationshipHandled
		}
	}
	return rel
}

// 处理喜欢不喜欢及匹配的关系
func HandleRelationship(rel po.Relationship) []po.Relationship {
	var rels []po.Relationship

	// 新增不喜欢时，直接返回 todo 常量待枚举
	if rel.State == "disliked" {
		rels = append(rels, rel)
		return rels
	}
	// 新增喜欢时，存在被喜欢关系
	relDao := new(dao.RelationshipDao)
	one, exist := relDao.GetOne(rel.OtherUserId, rel.UserId)
	if exist {
		if "liked" == one.State {
			// 关系转为喜欢
			one.State = "matched"
			rel.State = "matched"
			rels = append(rels, rel, one)
			return rels
		}
	}
	rels = append(rels, rel)
	return rels
}
