package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/textures1245/BlogDuaaeeg-backend/model/user-follower/entity"
)

type usrFollowerCon struct {
	UsrFollowerUse entity.UserFollowerService
}

func NewUsrFollowerController(usrFollowerUse entity.UserFollowerService) *usrFollowerCon {
	return &usrFollowerCon{
		usrFollowerUse,
	}
}

func (u *usrFollowerCon) SubscribeUser(c *gin.Context) {
	usrFollowerUuid := c.Param("user_uuid")
	if usrFollowerUuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "User UUID is required",
			"result":      nil,
		})
		return

	}

	req := new(entity.UserFollowerReqDat)
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
		return
	}

	res, err := u.UsrFollowerUse.OnSubscribeUser(usrFollowerUuid, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusText(http.StatusInternalServerError),
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

func (u *usrFollowerCon) UnsubscribeUser(c *gin.Context) {
	usrFollowerUuid := c.Param("user_uuid")
	if usrFollowerUuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "User UUID is required",
			"result":      nil,
		})
		return
	}

	req := new(entity.UserFollowerReqDat)
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
		return
	}

	err := u.UsrFollowerUse.OnUnsubscribeUser(usrFollowerUuid, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusText(http.StatusInternalServerError),
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
		"result":      nil,
	})
}
