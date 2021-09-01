package vo

// 接收的vo
type RelationshipInVo struct {
	UserId      int    `json:"user_id"`
	OtherUserId int    `json:"other_user_id"`
	State       string `json:"state" binding:"oneof=liked disliked"`
}
