package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/textures1245/BlogDuaaeeg-backend/model/auth/entity"
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
