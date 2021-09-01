package main

import (
	"fmt"
	"github.com/gin-gonic/examples/onboarding/dao"
	"github.com/gin-gonic/examples/onboarding/po"
	"github.com/gin-gonic/examples/onboarding/service"
	"time"
)

func testUserDao() {
	u := new(dao.UserDao)

	// 新增用户
	user := po.User{
		Name:       "探探用户",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Email:      "email"}

	user1 := po.User{
		Name:       "探探用户1",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Email:      "email1"}
	user2 := po.User{
		Name:       "探探用户2",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Email:      "email2"}
	user3 := po.User{
		Name:       "探探用户3",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Email:      "email3"}
	add := u.Add(user)
	add1 := u.Add(user1)
	add2 := u.Add(user2)
	add3 := u.Add(user3)
	fmt.Println("用户1插入前：", user, "，插入后：", add)
	fmt.Println("用户1插入前：", user1, "，插入后：", add1)
	fmt.Println("用户1插入前：", user2, "，插入后：", add2)
	fmt.Println("用户1插入前：", user3, "，插入后：", add3)
	// 查询所有用户
	users := u.QueryAll()
	fmt.Println(users)
}

func testUserService() {
	u := new(service.UserService)

	// 新增用户
	user := po.User{
		Name:       "探探用户5",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Email:      "email"}

	user1 := po.User{
		Name:       "探探用户6",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Email:      "email1"}
	user2 := po.User{
		Name:       "探探用户7",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Email:      "email2"}
	user3 := po.User{
		Name:       "探探用户8",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Email:      "email3"}
	add := u.AddUser(user)
	add1 := u.AddUser(user1)
	add2 := u.AddUser(user2)
	add3 := u.AddUser(user3)
	fmt.Println("用户1插入前：", user, "，插入后：", add)
	fmt.Println("用户1插入前：", user1, "，插入后：", add1)
	fmt.Println("用户1插入前：", user2, "，插入后：", add2)
	fmt.Println("用户1插入前：", user3, "，插入后：", add3)
	// 查询所有用户
	users := u.QueryAll()
	fmt.Println(users)
}

func testRelationship() {
	relService := new(service.RelationshipService)

	// 无关系时新增喜欢
	rel := po.Relationship{
		UserId:      43,
		OtherUserId: 44,
		State:       "liked",
		CreateTime:  time.Now(),
		UpdateTime:  time.Now()}
	// 新增不喜欢
	rel1 := po.Relationship{
		UserId:      44,
		OtherUserId: 45,
		State:       "disliked",
		CreateTime:  time.Now(),
		UpdateTime:  time.Now()}
	// 不喜欢时新增喜欢
	rel2 := po.Relationship{
		UserId:      45,
		OtherUserId: 44,
		State:       "liked",
		CreateTime:  time.Now(),
		UpdateTime:  time.Now()}
	// 喜欢时新增喜欢，两种关系需重置为匹配
	rel3 := po.Relationship{
		UserId:      44,
		OtherUserId: 43,
		State:       "liked",
		CreateTime:  time.Now(),
		UpdateTime:  time.Now()}

	add := relService.AddRelationship(rel)
	add1 := relService.AddRelationship(rel1)
	add2 := relService.AddRelationship(rel2)
	add3 := relService.AddRelationship(rel3)
	fmt.Println("关系1插入前：", rel, "，插入后：", add)
	fmt.Println("关系2插入前：", rel1, "，插入后：", add1)
	fmt.Println("关系3插入前：", rel2, "，插入后：", add2)
	fmt.Println("关系4插入前：", rel3, "，插入后：", add3)
	// 查询所有用户
	var queryUserId int = 44
	rels := relService.QueryAll(queryUserId)
	fmt.Println(rels)
}

func main() {
	testRelationship()
	//testUserService()
}
