package entities

type UserFollowerResDat struct {
	ID               string `json:"id"`
	UserFollowerUuid string `json:"user_follower_uuid"`
	UserFolloweeUuid string `json:"user_followee_uuid"`
	CreatedAt        string `json:"created_at"`
	UpdateAt         string `json:"update_at"`
}
