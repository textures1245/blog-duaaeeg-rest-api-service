package dtos

type PostCategoryReqDat struct {
	Name string `json:"name" db:"name" form:"name" binding:"required" validate:"required"`
}
