package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_errEntity "github.com/textures1245/BlogDuaaeeg-backend/api/error/entity"
	"github.com/textures1245/BlogDuaaeeg-backend/api/error/handle"
	"github.com/textures1245/BlogDuaaeeg-backend/api/model/auth/entity"
)

type authCon struct {
	AuthUse entity.AuthService
}

func NewAuthController(authService entity.AuthService) *authCon {
	controller := &authCon{
		AuthUse: authService,
	}

	return controller
}

func (h *authCon) Login(c *gin.Context) {
	req := new(entity.UsersCredentials)
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
		return
	}

	res, err := h.AuthUse.Login(req)
	if err != nil {
		handlerE := handle.NewHandler(&handle.HandleUse{})
		hE := handlerE.PrismaAuthHandle(*err.(*_errEntity.CError))
		c.JSON(hE.StatusCode, gin.H{
			"status":      http.StatusText(hE.StatusCode),
			"status_code": hE.StatusCode,
			"message":     hE.Error(),
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

func (h *authCon) Register(c *gin.Context) {
	req := new(entity.UsersCredentials)
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
		return
	}

	res, err := h.AuthUse.Register(req)
	if err != nil {
		handlerE := handle.NewHandler(&handle.HandleUse{})
		hE := handlerE.PrismaAuthHandle(*err.(*_errEntity.CError)) // Pass the value of cE instead of its pointer
		c.JSON(hE.StatusCode, gin.H{
			"status":      http.StatusText(hE.StatusCode),
			"status_code": hE.StatusCode,
			"message":     hE.Error(),
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

func (h *authCon) AuthTest(c *gin.Context) {
	uuid := c.MustGet("user_uuid")
	email := c.MustGet("email")

	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"status_code": http.StatusOK,
		"message":     "",
		"result": gin.H{
			"user_uuid": uuid,
			"email":     email,
		},
	})
}
