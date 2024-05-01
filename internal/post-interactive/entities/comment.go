package entities

type CommentResDat struct {
	UUID      string `json:"uuid"`
	UserUuid  string `json:"user_uuid"`
	Comment   string `json:"comment"`
	PostUUID  string `json:"post_uuid"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}
