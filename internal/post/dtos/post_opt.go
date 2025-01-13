package dtos

type FetchPostOptReq struct {
	Page  int `json:"page" form:"page" binding:"required" validate:"required"`
	Limit int `json:"limit" form:"limit" binding:"required" validate:"required"`
}
