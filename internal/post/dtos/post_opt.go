package dtos

type FetchPostOptReq struct {
	Page int `json:"page" form:"page" binding:"required" validate:"required"`
}
