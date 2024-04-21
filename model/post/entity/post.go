package entity

type PostReqDat struct {
	UserUuid  string `json:"user_uuid" db:"user_uuid" form:"user_uuid" binding:"required" validate:"required"`
	Title     string `json:"title" db:"title" form:"title" binding:"required" validate:"required"`
	Content   string `json:"content" db:"content" form:"content" binding:"required" validate:"required"`
	Published bool   `json:"published" db:"published" form:"published" binding:"required" validate:"required"`
}
