package dao

import (
	"fmt"
	"github.com/gin-gonic/examples/onboarding/po"
	"time"
)

type UserDao struct {
	Id int
}

func (u UserDao) Add(user po.User) po.User {
	db := getDb()
	defer db.Close()
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	_, err := db.Model(&user).Insert()
	PanicIf(err)
	return user
}

func (u UserDao) QueryAll() []po.User {
	db := getDb()
	defer db.Close()
	var users []po.User
	err := db.Model(&users).Select()
	PanicIf(err)
	return users
}

var _ userDaoI = (*UserDao)(nil)

func main() {
	u := new(UserDao)
	users := u.QueryAll()
	fmt.Println(users)
}
