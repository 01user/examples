package dao

import (
	"github.com/gin-gonic/examples/onboarding/po"
	"time"
)

type RelationshipDao struct {
	Id int
}

func (relDao RelationshipDao) AddAll(rels []po.Relationship) []po.Relationship {
	db := getDb()
	defer db.Close()
	relationships := defaultTime(rels)
	_, err := db.Model(&relationships).
		OnConflict("(user_id,other_user_id) DO UPDATE").
		Set("state = EXCLUDED.state").
		Insert()
	PanicIf(err)
	return relationships
}

func defaultTime(rels []po.Relationship) []po.Relationship {
	var results []po.Relationship
	for _, rel := range rels {
		result := new(po.Relationship)
		result.CreateTime = time.Now()
		result.UpdateTime = time.Now()
		result.Id = rel.Id
		result.UserId = rel.UserId
		result.OtherUserId = rel.OtherUserId
		result.State = rel.State
		result.Type = rel.Type
		results = append(results, *result)
	}
	return results
}

func (relDao RelationshipDao) GetOne(userId int, otherUserId int) (po.Relationship, bool) {
	var rel po.Relationship
	db := getDb()
	defer db.Close()
	err := db.Model(&rel).
		Where("user_id = ? and other_user_id = ?", userId, otherUserId).
		First()
	if err != nil {
		return rel, false
	}
	return rel, true
}

func (relDao RelationshipDao) Add(rel po.Relationship) po.Relationship {
	db := getDb()
	defer db.Close()
	_, err := db.Model(&rel).Insert()
	PanicIf(err)
	return rel
}

func (relDao RelationshipDao) QueryAll(userId int) []po.Relationship {
	db := getDb()
	defer db.Close()
	var relationships []po.Relationship
	err := db.Model(&relationships).
		Where("user_id = ?", userId).
		Select()
	PanicIf(err)
	return relationships
}

var _ relationshipDaoI = (*RelationshipDao)(nil)
