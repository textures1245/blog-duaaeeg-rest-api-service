package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type postInterCtrl struct {
	PostInter entity.UserInteractiveService
}

func NewPostInteractiveController(PostInter entity.UserInteractiveService) *postInterCtrl {
	return &postInterCtrl{
		PostInter,
	}
}

func (h *postInterCtrl) UserCommentedToPost(c *gin.Context) {
	req := new(entity.CommentReqDat)
	pUuid := c.Param("post_uuid")
	if pUuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "Post UUID is required",
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

	res, err := h.UsrInter.OnCreateNewComment(pUuid, req)
	if err != nil {
		customErrorHandle("CommentModel", c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"status_code": http.StatusOK,
		"message":     "",
		"result":      res,
	})
}

func (h *postInterCtrl) UserUpdateComment(c *gin.Context) {
	cUuid := c.Param("comment_uuid")
	if cUuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "Comment UUID is required",
			"result":      nil,
		})
		return
	}

	req := new(entity.CommentReqDat)
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
		return
	}

	res, err := h.UsrInter.OnUpdateCommentByUUID(cUuid, req)
	if err != nil {
		customErrorHandle("CommentModel", c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"status_code": http.StatusOK,
		"message":     "",
		"result":      res,
	})
}

func (h *postInterCtrl) UserDeleteComment(c *gin.Context) {
	cUuid := c.Param("comment_uuid")
	if cUuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "Comment UUID is required",
			"result":      nil,
		})
		return
	}

	err := h.UsrInter.OnDeleteCommentByUUID(cUuid)
	if err != nil {
		customErrorHandle("CommentModel", c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"status_code": http.StatusOK,
		"message":     "Comment has been deleted",
		"result":      "",
	})
}

func (h *postInterCtrl) UserLikedPost(c *gin.Context) {
	pUuid := c.Param("post_uuid")
	if pUuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "Post UUID is required",
			"result":      nil,
		})
		return
	}

	req := new(entity.LikeReqDat)
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
		return
	}

	res, err := h.UsrInter.OnLikedPost(pUuid, req)
	if err != nil {
		customErrorHandle("LikeModel", c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"status_code": http.StatusOK,
		"message":     "",
		"result":      res,
	})
}

func (h *postInterCtrl) UserUnlikedPost(c *gin.Context) {
	pUuid := c.Param("post_uuid")
	if pUuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "Post UUID is required",
			"result":      nil,
		})
		return
	}

	req := new(entity.LikeReqDat)
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
		return
	}

	err := h.UsrInter.OnUnlikedPost(pUuid, req)
	if err != nil {
		customErrorHandle("LikeModel", c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"status_code": http.StatusOK,
		"message":     "Post has been unliked",
		"result":      "",
	})
}

func customErrorHandle(nameM string, c *gin.Context, err error) {
	handlerE := handler.NewHandler(&handler.HandleUse{})
	hE := handlerE.PrismaCustomHandle(nameM, *err.(*_errEntity.CError))
	c.JSON(hE.StatusCode, gin.H{
		"status":      http.StatusText(hE.StatusCode),
		"status_code": hE.StatusCode,
		"message":     hE.Error(),
		"result":      nil,
	})
}
