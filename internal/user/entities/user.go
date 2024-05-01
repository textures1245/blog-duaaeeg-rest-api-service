package entities

import "github.com/textures1245/BlogDuaaeeg-backend/db"

type UserResDat struct {
	UUID        string                 `json:"uuid"`
	Email       string                 `json:"email"`
	UserProfile *UserProfileRes        `json:"user_profile"`
	Subscribers []db.UserFollowerModel `json:"subscriber" db:"subscriber"`
	Subscribing []db.UserFollowerModel `json:"subscribing" db:"subscribing"`
}
