package entities

import (
	"github.com/textures1245/BlogDuaaeeg-backend/db"
	entityCate "github.com/textures1245/BlogDuaaeeg-backend/internal/category/entities"
	entityUser "github.com/textures1245/BlogDuaaeeg-backend/internal/user/entities"
)

// TODO: Remove UserUUID from PostReqDat, to refactor data req structor

type PostResDat struct {
	UUID              string                         `json:"uuid"`
	User              entityUser.UserResDat          `json:"user"`
	Title             string                         `json:"title"`
	ImgBannerUrl      *string                        `json:"img_banner_url"`
	Source            string                         `json:"source"`
	Published         bool                           `json:"published"`
	SrcType           string                         `json:"src_type"`
	PublishedPostUUID string                         `json:"published_post_uuid"`
	PostUUID          string                         `json:"post_uuid"`
	Category          *entityCate.PostCategoryResDat `json:"category"`
	Tags              *entityCate.PostTagResDat      `json:"tags"`
	Comments          []db.CommentModel              `json:"comments"`
	Like              []db.LikeModel                 `json:"like"`
	CreatedAt         string                         `json:"created_at"`
	UpdateAt          string                         `json:"update_at"`
}

type PostWithTagCateResDat struct {
	Post     *PostResDat                    `json:"post"`
	Category *entityCate.PostCategoryResDat `json:"category"`
	Tags     *entityCate.PostTagResDat      `json:"tags"`
}
