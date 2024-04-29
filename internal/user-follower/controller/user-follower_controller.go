package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_errEntity "github.com/textures1245/BlogDuaaeeg-backend/error/entity"
	"github.com/textures1245/BlogDuaaeeg-backend/error/handler"
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
		usrFollowerHandle(c, err)
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
		usrFollowerHandle(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"status_code": http.StatusOK,
		"message":     "",
		"result":      nil,
	})
}

func usrFollowerHandle(c *gin.Context, err error) {
	handle := handler.NewHandler(&handler.HandleUse{})
	hE := handle.PrismaCustomHandle("UserFollowerModel", *err.(*_errEntity.CError))
	c.JSON(hE.StatusCode, gin.H{
		"status":      http.StatusText(hE.StatusCode),
		"status_code": hE.StatusCode,
		"message":     hE.Error(),
		"result":      nil,
	})
}
