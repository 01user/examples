package vo

// 吐出的vo
type RelationshipVo struct {
	OtherUserId int    `json:"user_id"`
	State       string `json:"state"`
	Type        string `json:"type"`
}
