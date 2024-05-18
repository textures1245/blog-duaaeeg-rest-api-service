package dtos

type UserProfileDataRequest struct {
	FirstName      string `json:"first_name" db:"first_name" form:"first_name" binding:"required" validate:"required"`
	LastName       string `json:"last_name" db:"last_name" form:"last_name" binding:"required" validate:"required"`
	Bio            string `json:"bio" db:"bio" form:"bio" binding:"required" validate:"required"`
	ProfilePicture string `json:"profile_picture" db:"profile_pic" form:"profile_picture" binding:"required" validate:"required"`
}
