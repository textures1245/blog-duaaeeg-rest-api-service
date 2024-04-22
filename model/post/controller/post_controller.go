package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_errEntity "github.com/textures1245/BlogDuaaeeg-backend/error/entity"
	"github.com/textures1245/BlogDuaaeeg-backend/error/handler"
	cateEntity "github.com/textures1245/BlogDuaaeeg-backend/model/category/entity"
	"github.com/textures1245/BlogDuaaeeg-backend/model/post/entity"
)

type postCon struct {
	PostUse entity.PostService
	CateUse cateEntity.PostTagCateService
}

func NewPostController(PostUse entity.PostService, CateUse cateEntity.PostTagCateService) *postCon {
	return &postCon{
		PostUse,
		CateUse,
	}

}

func (h *postCon) CreatePost(c *gin.Context) {
	req := new(entity.PostReqDat)
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
		return
	}

	res, err := h.PostUse.OnCreateNewPost(req)
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
	postCate, err := h.CateUse.OnCreateCategory(res.UUID, req.PostCategory)
	if err != nil {
		postErrorHandle(c, err)
		return
	}
	postTag, err := h.CateUse.OnCreateTags(res.UUID, req.PostTag)
	if err != nil {
		postErrorHandle(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"status_code": http.StatusOK,
		"message":     "",
		"result": &entity.PostWithTagCateResDat{
			Post:     res,
			Category: postCate,
			Tags:     postTag,
		},
	})
}

func (h *postCon) UpdatePost(c *gin.Context) {
	req := new(entity.PostReqDat)
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

	res, err := h.PostUse.OnUpdatePostByUUID(pUuid, req)
	if err != nil {
		postErrorHandle(c, err)
		return
	}

	postCate, err := h.CateUse.OnUpdateCategory(res.Category.ID, req.PostCategory)
	if err != nil {
		postErrorHandle(c, err)
		return
	}
	postTag, err := h.CateUse.OnUpdateTags(res.Tags.ID, req.PostTag)
	if err != nil {
		postErrorHandle(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"status_code": http.StatusOK,
		"message":     "",
		"result": &entity.PostWithTagCateResDat{
			Post:     res,
			Category: postCate,
			Tags:     postTag,
		},
	})
}

func (h *postCon) GetPostByUUID(c *gin.Context) {
	pUuid := c.Param("post_uuid")

	res, err := h.PostUse.OnFetchPostByUUID(pUuid)
	if err != nil {
		postErrorHandle(c, err)
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
		postErrorHandle(c, err)
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

	res, err := h.PostUse.OnFetchPublisherPosts(&entity.FetchPostOptReq{
		Page: page,
	})
	if err != nil {
		postErrorHandle(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"status_code": http.StatusOK,
		"message":     "",
		"result":      res,
	})
}

func postErrorHandle(c *gin.Context, err error) {
	handlerE := handler.NewHandler(&handler.HandleUse{})
	hE := handlerE.PrismaPostHandle(*err.(*_errEntity.CError))
	c.JSON(hE.StatusCode, gin.H{
		"status":      http.StatusText(hE.StatusCode),
		"status_code": hE.StatusCode,
		"message":     hE.Error(),
		"result":      nil,
	})
}
