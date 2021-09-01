package dao

import (
	"fmt"
	"github.com/gin-gonic/examples/onboarding/po"
	"time"
)

func batchInsert() {
	db := getDb()
	user1 := &po.User{Email: "email1-1",
		UserName:   "用户1-1",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		//DeleteTime: time.Now(),
	}
	user2 := &po.User{Email: "email2-1",
		UserName:   "用户2-2",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		//DeleteTime: time.Now(),
	}
	users := []po.User{{
		Email:      "email3-3",
		UserName:   "用户3-3",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		//DeleteTime: time.Now(),
	}, {
		Email:      "email4-4",
		UserName:   "用户4-4",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		//DeleteTime: time.Now(),
	}}
	_, err := db.Model(user1, user2).Insert()
	PanicIf(err)
	_, err = db.Model(&users).Insert()
	PanicIf(err)
	fmt.Println(user1, user2, users)
}
