package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_errEntity "github.com/textures1245/BlogDuaaeeg-backend/error/entity"
	"github.com/textures1245/BlogDuaaeeg-backend/model/user/entity"
)

type userCon struct {
	UserUse entity.UserService
}

func NewUserController(userService entity.UserService) *userCon {
	return &userCon{
		UserUse: userService,
	}
}

func (h *userCon) UpdateUserProfile(c *gin.Context) {
	req := new(entity.UserProfileDataRequest)
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
		return
	}

	res, err := h.UserUse.OnUpdateUserProfile(req)
	if err != nil {
		cE := err.(*_errEntity.CError)
		c.JSON(cE.StatusCode, gin.H{
			"status":      http.StatusText(cE.StatusCode),
			"status_code": http.StatusInternalServerError,
			"message":     err.Error(),
			"result":      nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"status_code": http.StatusOK,
		"message":     "",
		"result":      res,
	})
}
