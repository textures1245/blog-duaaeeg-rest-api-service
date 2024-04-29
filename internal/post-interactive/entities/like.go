package entities

type LikeResDat struct {
	UUID      string `json:"uuid"`
	UserUuid  string `json:"user_uuid"`
	PostUuid  string `json:"post_uuid"`
	CreatedAt string `json:"created_at"`
}
