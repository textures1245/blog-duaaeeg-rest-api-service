package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/user"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/user/dtos"
	_errEntity "github.com/textures1245/BlogDuaaeeg-backend/pkg/error/entity"
	"github.com/textures1245/BlogDuaaeeg-backend/pkg/error/handler"
)

type userCon struct {
	UserUse user.UserService
}

func NewUserController(userService user.UserService) *userCon {
	return &userCon{
		UserUse: userService,
	}
}

func (h *userCon) FetchUserByUUID(c *gin.Context) {
	userUuid := c.Param("user_uuid")
	if userUuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "User UUID is required",
			"result":      nil,
		})
		return
	}

	res, err := h.UserUse.OnFetchUserByUUID(userUuid)
	if err != nil {
		handlerE := handler.NewHandler(&handler.HandleUse{})
		hE := handlerE.PrismaCustomHandle("UserModel", *err.(*_errEntity.CError))
		c.JSON(hE.StatusCode, gin.H{
			"status":      http.StatusText(hE.StatusCode),
			"status_code": hE.StatusCode,
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

func (h *userCon) UpdateUserProfile(c *gin.Context) {
	req := new(dtos.UserProfileDataRequest)

	userUuid := c.Param("user_uuid")
	if userUuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "User UUID is required",
			"result":      nil,
		})
		return
	}

	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
		return
	}

	res, err := h.UserUse.OnUpdateUserProfile(userUuid, req)
	if err != nil {
		handlerE := handler.NewHandler(&handler.HandleUse{})
		hE := handlerE.PrismaAuthHandle(*err.(*_errEntity.CError))
		c.JSON(hE.StatusCode, gin.H{
			"status":      http.StatusText(hE.StatusCode),
			"status_code": hE.StatusCode,
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

func (h *userCon) FetchUsers(c *gin.Context) {
	fetchMode := c.Request.Header.Get("fetch-mode")
	if fetchMode == "WITH_PASSWORD" {
		res, err := h.UserUse.OnFetchUsersWithPW()
		if err != nil {
			handlerE := handler.NewHandler(&handler.HandleUse{})
			hE := handlerE.PrismaCustomHandle("UserModel", *err.(*_errEntity.CError))
			c.JSON(hE.StatusCode, gin.H{
				"status":      http.StatusText(hE.StatusCode),
				"status_code": hE.StatusCode,
				"message":     err.Error(),
				"result":      nil,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"status":      "OK",
			"status_code": http.StatusOK,
			"message":     "",
			"result":      res,
		})
	} else {
		res, err := h.UserUse.OnFetchUsers()
		if err != nil {
			handlerE := handler.NewHandler(&handler.HandleUse{})
			hE := handlerE.PrismaCustomHandle("UserModel", *err.(*_errEntity.CError))
			c.JSON(hE.StatusCode, gin.H{
				"status":      http.StatusText(hE.StatusCode),
				"status_code": hE.StatusCode,
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

}

func (h *userCon) DeleteUser(c *gin.Context) {
	userUuid := c.Param("user_uuid")
	if userUuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "User UUID is required",
			"result":      nil,
		})
		return
	}

	err := h.UserUse.OnDeleteUser(userUuid)
	if err != nil {
		handlerE := handler.NewHandler(&handler.HandleUse{})
		hE := handlerE.PrismaCustomHandle("UserModel", *err.(*_errEntity.CError))
		c.JSON(hE.StatusCode, gin.H{
			"status":      http.StatusText(hE.StatusCode),
			"status_code": hE.StatusCode,
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
