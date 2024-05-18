package entities

import "github.com/textures1245/BlogDuaaeeg-backend/db"

type UserResDat struct {
	UUID        string                 `json:"uuid"`
	Email       string                 `json:"email"`
	UserProfile *UserProfileRes        `json:"user_profile"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
	Subscribers []db.UserFollowerModel `json:"subscriber" db:"subscriber"`
	Subscribing []db.UserFollowerModel `json:"subscribing" db:"subscribing"`
}
type UserWithPWResDat struct {
	UUID        string                 `json:"uuid"`
	Email       string                 `json:"email"`
	Password    string                 `json:"password"`
	UserProfile *UserProfileRes        `json:"user_profile"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
	Subscribers []db.UserFollowerModel `json:"subscriber" db:"subscriber"`
	Subscribing []db.UserFollowerModel `json:"subscribing" db:"subscribing"`
}
