package po

import (
	"time"
)

type User struct {
	Id         int
	Name       string
	Type       string
	UserName   string
	Email      string
	CreateTime time.Time
	UpdateTime time.Time
	DeleteTime time.Time
}
