package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/auth"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/auth/entities"
	_errEntity "github.com/textures1245/BlogDuaaeeg-backend/pkg/error/entity"
	"github.com/textures1245/BlogDuaaeeg-backend/pkg/error/handler"
)

type authCon struct {
	AuthUse auth.AuthUsecase
}

func NewAuthController(authService auth.AuthUsecase) *authCon {
	controller := &authCon{
		AuthUse: authService,
	}

	return controller
}

func (h *authCon) Login(c *gin.Context) {
	req := new(entities.UsersCredentials)
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
		handlerE := handler.NewHandler(&handler.HandleUse{})
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
	req := new(entities.UsersCredentials)
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
		handlerE := handler.NewHandler(&handler.HandleUse{})
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
