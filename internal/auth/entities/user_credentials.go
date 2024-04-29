package entities

type UsersCredentials struct {
	Email    string `json:"email" db:"email" form:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" db:"password" form:"password" binding:"required" validate:"required,min=5"`
}
