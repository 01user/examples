package po

import (
	"time"
)

type Relationship struct {
	Id          int
	UserId      int
	OtherUserId int
	State       string
	Type        string
	CreateTime  time.Time
	UpdateTime  time.Time
	DeleteTime  time.Time
}
