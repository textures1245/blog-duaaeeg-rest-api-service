package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/category"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/post"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/post/dtos"
	_errEntity "github.com/textures1245/BlogDuaaeeg-backend/pkg/error/entity"
	"github.com/textures1245/BlogDuaaeeg-backend/pkg/error/handler"
)

type postCon struct {
	PostUse post.PostService
	CateUse category.PostTagCateService
}

func NewPostController(PostUse post.PostService, CateUse category.PostTagCateService) *postCon {
	return &postCon{
		PostUse,
		CateUse,
	}

}

func (h *postCon) CreatePost(c *gin.Context) {
	req := new(dtos.PostReqDat)
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
		return
	}

	postCate, err := h.CateUse.OnCreateOrUpdateCategory(req.PostCategory)
	if err != nil {
		customErrorHandle("PostModel", c, err)
		return
	}
	postTag, err := h.CateUse.OnCreateTags(req.PostTag)
	if err != nil {
		customErrorHandle("PostModel", c, err)
		return
	}

	res, err := h.PostUse.OnCreateNewPost(c, postCate, postTag, req)
	if err != nil {
		handlerE := handler.NewHandler(&handler.HandleUse{})
		hE := handlerE.PrismaPostHandle(*err.(*_errEntity.CError))
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

func (h *postCon) UpdatePost(c *gin.Context) {
	req := new(dtos.PostReqDat)
	pUuid := c.Param("post_uuid")

	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
		return
	}

	postCate, err := h.CateUse.OnCreateOrUpdateCategory(req.PostCategory)
	if err != nil {
		customErrorHandle("PostModel", c, err)
		return
	}

	res, err := h.PostUse.OnUpdatePostAndTagByUUID(c, postCate, pUuid, req)
	if err != nil {
		customErrorHandle("PostModel", c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"status_code": http.StatusOK,
		"message":     "",
		"result":      res,
	})
}

func (h *postCon) GetPostByUUID(c *gin.Context) {
	pUuid := c.Param("post_uuid")

	res, err := h.PostUse.OnFetchPostByUUID(pUuid)
	if err != nil {
		customErrorHandle("PostModel", c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"status_code": http.StatusOK,
		"message":     "",
		"result":      res,
	})
}

func (h *postCon) GetPostByUserUUID(c *gin.Context) {
	uUuid := c.Param("user_uuid")

	res, err := h.PostUse.OnFetchOwnerPosts(uUuid)
	if err != nil {
		customErrorHandle("PostModel", c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"status_code": http.StatusOK,
		"message":     "",
		"result":      res,
	})
}

func (h *postCon) GetPublisherPosts(c *gin.Context) {
	p := c.Query("page")

	if p == "" {
		p = "0"
	}

	page, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "Page query params is invalid",
			"result":      nil,
		})
		return
	}

	res, err := h.PostUse.OnFetchPublisherPosts(&dtos.FetchPostOptReq{
		Page: page,
	})
	if err != nil {
		customErrorHandle("PostModel", c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"status_code": http.StatusOK,
		"message":     "",
		"result":      res,
	})
}

func (h *postCon) DeletePostAndPublisherPostByUUID(c *gin.Context) {
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

	err := h.PostUse.OnDeletePostByUUID(pUuid)
	if err != nil {
		customErrorHandle("PostModel", c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"status_code": http.StatusOK,
		"message":     "Post and PublisherPost has been deleted",
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
