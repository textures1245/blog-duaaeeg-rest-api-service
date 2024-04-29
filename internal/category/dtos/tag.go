package dtos

type PostTagReqDat struct {
	Tags []string `json:"tags" db:"tag" form:"tag" binding:"required" validate:"required"`
}
