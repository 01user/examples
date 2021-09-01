package dao

import (
	"fmt"
	"github.com/gin-gonic/examples/onboarding/po"
	"time"
)

func batchUpdate() {
	db := getDb()
	user1 := &po.User{
		Id:         23,
		Email:      "email1更改邮箱",
		UserName:   "用户1",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		//DeleteTime: time.Now(),
	}
	user2 := &po.User{
		Id:         24,
		Email:      "email2更改邮箱",
		UserName:   "用户2",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		//DeleteTime: time.Now(),
	}
	users := []po.User{{
		Id:         25,
		Email:      "email3更改邮箱",
		UserName:   "用户3",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		//DeleteTime: time.Now(),
	}, {
		Id:         26,
		Email:      "email4更改邮箱",
		UserName:   "用户4",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		//DeleteTime: time.Now(),
	}}
	_, err := db.Model(user1, user2).
		// 需修改的列
		Column("email", "update_time").
		// 根据主键条件进行修改
		Update()
	PanicIf(err)
	_, err = db.Model(&users). // 需修改的列
					Column("email", "update_time").Update()
	PanicIf(err)
	fmt.Println(user1, user2, users)
	var users1 []po.User
	err = db.Model(&users1).Order("id").Select()
	if err != nil {
		panic(err)
	}
	fmt.Println(users1)

	var user5 po.User
	users2, err := db.Model(&user5).
		// 根据条件进行修改
		Set("email = ?", "修改后的邮箱").
		Where("user_name = ?", "用户3").
		Returning("*").
		Update()
	PanicIf(err)
	fmt.Println(users2)
}
