package usecase

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/textures1245/BlogDuaaeeg-backend/db"
	cate "github.com/textures1245/BlogDuaaeeg-backend/internal/category"
	cateEntities "github.com/textures1245/BlogDuaaeeg-backend/internal/category/entities"
	entityEntities "github.com/textures1245/BlogDuaaeeg-backend/internal/category/entities"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/file"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/file/entities"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/post"
	postDtos "github.com/textures1245/BlogDuaaeeg-backend/internal/post/dtos"
	postEntities "github.com/textures1245/BlogDuaaeeg-backend/internal/post/entities"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/user"
	userEntities "github.com/textures1245/BlogDuaaeeg-backend/internal/user/entities"
)

type postUse struct {
	PostRepo  post.PostRepository
	UsersRepo user.UsersRepository
	TagRepo   cate.PostTagRepository
	fileUse   file.FileUsecase
}

// OnFetchPublisherPostByUUID implements post.PostService.
func (u *postUse) OnFetchPublisherPostByUUID(pbUuid string) (*postEntities.PostResDat, error) {
	pbPost, err := u.PostRepo.FetchPublisherPostByUUID(pbUuid)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	res := &postEntities.PostResDat{
		UUID: pbPost.PostUUID,
		User: userEntities.UserResDat{
			UUID:  pbPost.Post().User().UUID,
			Email: pbPost.Post().User().Email,
			UserProfile: func() *userEntities.UserProfileRes {
				if v, ok := pbPost.Post().User().UserProfile(); ok {
					return &userEntities.UserProfileRes{
						FirstName:      v.FirstName,
						LastName:       v.LastName,
						ProfilePicture: v.ProfilePicture,
					}
				}
				return nil
			}(),
		},
		ImgBannerUrl: func() *string {
			if v, ok := pbPost.Post().ImageBannerURL(); ok {
				return &v
			}
			return nil
		}(),
		Title:             pbPost.Post().Title,
		Source:            pbPost.Post().Source,
		Published:         pbPost.Post().Published,
		SrcType:           string(pbPost.Post().SrcType),
		PublishedPostUUID: pbPost.Post().UUID,
		Category: &cateEntities.PostCategoryResDat{
			ID:   pbPost.Post().Category().ID,
			Name: pbPost.Post().Category().Name,
		},
		Tags: &cateEntities.PostTagResDat{
			ID:   pbPost.Post().Tags().ID,
			Tags: pbPost.Post().Tags().Tags,
		},
		CreatedAt: pbPost.Post().CreatedAt.String(),
		UpdateAt:  pbPost.Post().UpdatedAt.String(),
		Comments:  pbPost.Post().Comments(),
		Like:      pbPost.Post().Likes(),
	}

	return res, nil
}

func NewPostService(postRepo post.PostRepository, usersRepo user.UsersRepository, tagRepo cate.PostTagRepository, fileUse file.FileUsecase) post.PostService {
	return &postUse{
		PostRepo:  postRepo,
		UsersRepo: usersRepo,
		TagRepo:   tagRepo,
		fileUse:   fileUse,
	}
}

func (u *postUse) OnCreateNewPost(c *gin.Context, cateResDat *entityEntities.PostCategoryResDat, tagResDat *entityEntities.PostTagResDat, req *postDtos.PostReqDat) (*postEntities.PostResDat, error) {
	ctx := c.Request.Context()

	if req.SrcType != "CONTENT" {
		file, _, err := u.fileUse.OnUploadFile(c, ctx, &entities.FileUploaderReq{
			FileName: req.Title,
			FileData: req.Content,
			FileType: req.SrcType,
		})
		if err != nil {
			log.Error(err)
			return nil, err
		}
		req.Content = file.FileData
	}

	post, err := u.PostRepo.CreatePost(cateResDat, tagResDat, req)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// check if post marked as publish then link to publication post
	pbpUuid := ""
	if oldUuid, isNil := post.PublishPostUUID(); !isNil {
		if post.Published {
			uuid, err := u.PostRepo.UpdatePostToPublisher(post.UserUUID, post.UUID)
			if err != nil {
				log.Error(err)
				log.Error(err)
				return nil, err
			}
			pbpUuid = uuid
		}
	} else {
		pbpUuid = oldUuid
	}

	res := &postEntities.PostResDat{
		UUID: post.UUID,
		User: userEntities.UserResDat{
			UUID:  post.User().UUID,
			Email: post.User().Email,
			UserProfile: func() *userEntities.UserProfileRes {
				if v, ok := post.User().UserProfile(); ok {
					return &userEntities.UserProfileRes{
						FirstName:      v.FirstName,
						LastName:       v.LastName,
						ProfilePicture: v.ProfilePicture,
					}
				}
				return nil
			}(),
		},
		Title:             post.Title,
		Source:            post.Source,
		Published:         post.Published,
		SrcType:           string(post.SrcType),
		PublishedPostUUID: pbpUuid,
		Category: &cateEntities.PostCategoryResDat{
			ID:   post.Category().ID,
			Name: post.Category().Name,
		},
		Tags: &cateEntities.PostTagResDat{
			ID:   post.Tags().ID,
			Tags: post.Tags().Tags,
		},
		CreatedAt: post.CreatedAt.String(),
		UpdateAt:  post.UpdatedAt.String(),
		Comments:  post.Comments(),
		Like:      post.Likes(),
	}
	return res, nil

}

func (u *postUse) OnUpdatePostAndTagByUUID(c *gin.Context, cateResDat *cateEntities.PostCategoryResDat, uuid string, req *postDtos.PostReqDat) (*postEntities.PostResDat, error) {
	ctx := c.Request.Context()

	file, _, err := u.fileUse.OnUploadFile(c, ctx, &entities.FileUploaderReq{
		FileName: req.Title,
		FileData: req.Content,
		FileType: req.SrcType,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	req.Content = file.FileData

	post, err := u.PostRepo.UpdatePostByUUID(cateResDat, uuid, req)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// check if post marked as publish then link to publication post
	pbpUuid := ""
	if oldUuid, isNil := post.PublishPostUUID(); !isNil {
		if post.Published {
			uuid, err := u.PostRepo.UpdatePostToPublisher(post.UserUUID, post.UUID)
			if err != nil {
				log.Error(err)
				return nil, err
			}
			pbpUuid = uuid
		}
	} else {
		pbpUuid = oldUuid
	}

	tagUpdated, err := u.TagRepo.UpdateTags(cateResDat.ID, req.PostTag)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	res := &postEntities.PostResDat{
		UUID: post.UUID,
		User: userEntities.UserResDat{
			UUID:  post.User().UUID,
			Email: post.User().Email,
			UserProfile: func() *userEntities.UserProfileRes {
				if v, ok := post.User().UserProfile(); ok {
					return &userEntities.UserProfileRes{
						FirstName:      v.FirstName,
						LastName:       v.LastName,
						ProfilePicture: v.ProfilePicture,
					}
				}
				return nil
			}(),
		},
		Title:             post.Title,
		Source:            post.Source,
		Published:         post.Published,
		SrcType:           string(post.SrcType),
		PublishedPostUUID: pbpUuid,
		Category: &cateEntities.PostCategoryResDat{
			ID:   post.Category().ID,
			Name: post.Category().Name,
		},
		Tags: &cateEntities.PostTagResDat{
			ID:   tagUpdated.ID,
			Tags: tagUpdated.Tags,
		},
		Comments:  post.Comments(),
		Like:      post.Likes(),
		CreatedAt: post.CreatedAt.String(),
		UpdateAt:  post.UpdatedAt.String(),
	}
	return res, nil

}

func (u *postUse) OnFetchPostByUUID(uuid string) (*postEntities.PostResDat, error) {
	post, err := u.PostRepo.FetchPostByUUID(uuid)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// pbpUuid, cateM, tagM := prismaOptKeyRetrieve(post)

	res := &postEntities.PostResDat{
		UUID: post.UUID,
		User: userEntities.UserResDat{
			UUID:  post.User().UUID,
			Email: post.User().Email,
			UserProfile: func() *userEntities.UserProfileRes {
				if v, ok := post.User().UserProfile(); ok {
					return &userEntities.UserProfileRes{
						FirstName:      v.FirstName,
						LastName:       v.LastName,
						ProfilePicture: v.ProfilePicture,
					}
				}
				return nil
			}(),
		},
		Title:             post.Title,
		Source:            post.Source,
		Published:         post.Published,
		SrcType:           string(post.SrcType),
		PublishedPostUUID: "",
		Category: &cateEntities.PostCategoryResDat{
			ID:   post.Category().ID,
			Name: post.Category().Name,
		},
		Tags: &cateEntities.PostTagResDat{
			ID:   post.Tags().ID,
			Tags: post.Tags().Tags,
		},
		Comments:  post.Comments(),
		Like:      post.Likes(),
		CreatedAt: post.CreatedAt.String(),
		UpdateAt:  post.UpdatedAt.String(),
	}
	if pbp, ok := post.PublishPostUUID(); ok {
		res.PublishedPostUUID = pbp
	}
	return res, nil
}

func (u *postUse) OnFetchOwnerPosts(userUuid string) ([]*postEntities.PostResDat, error) {
	posts, err := u.PostRepo.FetchPostByUserUUID(userUuid)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var res []*postEntities.PostResDat
	res = mapPostsDatToRes(posts, res)
	return res, nil

}

func (u *postUse) OnFetchPublisherPosts(opts *postDtos.FetchPostOptReq) ([]*postEntities.PostResDat, error) {
	posts, err := u.PostRepo.FetchPublisherPosts(opts)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var res []*postEntities.PostResDat
	for _, post := range posts {

		res = append(res, &postEntities.PostResDat{
			UUID: post.UUID,
			User: userEntities.UserResDat{
				UUID:  post.User().UUID,
				Email: post.User().Email,
				UserProfile: func() *userEntities.UserProfileRes {
					if v, ok := post.User().UserProfile(); ok {
						return &userEntities.UserProfileRes{
							FirstName:      v.FirstName,
							LastName:       v.LastName,
							ProfilePicture: v.ProfilePicture,
						}
					}
					return nil
				}(),
			},
			Title:  post.Post().Title,
			Source: post.Post().Source,
			ImgBannerUrl: func() *string {
				if v, ok := post.Post().ImageBannerURL(); ok {
					return &v
				}
				return nil
			}(),
			Published: post.Post().Published,
			SrcType:   string(post.Post().SrcType),
			PostUUID:  post.PostUUID,
			Category: &cateEntities.PostCategoryResDat{
				ID:   post.Post().Category().ID,
				Name: post.Post().Category().Name,
			},
			Tags: &cateEntities.PostTagResDat{
				ID:   post.Post().Tags().ID,
				Tags: post.Post().Tags().Tags,
			},
			Comments:  post.Post().Comments(),
			Like:      post.Post().Likes(),
			CreatedAt: post.CreatedAt.String(),
			UpdateAt:  post.UpdatedAt.String(),
		})
	}
	return res, nil
}

// - this func should be called as private func
func (u *postUse) OnSubmitPostToPublisher(userUuid string, postUuid string) (string, error) {
	pbpUuid, err := u.PostRepo.UpdatePostToPublisher(userUuid, postUuid)
	if err != nil {
		log.Error(err)
		return "", err
	}
	return pbpUuid, nil
}

func (u *postUse) OnDeletePostByUUID(postUuid string) error {

	err := u.PostRepo.DeletePostByUUID(postUuid)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil

}

func mapPostsDatToRes(pDat []db.PostModel, pRes []*postEntities.PostResDat) []*postEntities.PostResDat {
	for _, post := range pDat {

		rp := &postEntities.PostResDat{
			UUID: post.UUID,
			User: userEntities.UserResDat{
				UUID:  post.User().UUID,
				Email: post.User().Email,
				UserProfile: func() *userEntities.UserProfileRes {
					if v, ok := post.User().UserProfile(); ok {
						return &userEntities.UserProfileRes{
							FirstName:      v.FirstName,
							LastName:       v.LastName,
							ProfilePicture: v.ProfilePicture,
						}
					}
					return nil
				}(),
			},
			Title:             post.Title,
			Source:            post.Source,
			Published:         post.Published,
			SrcType:           string(post.SrcType),
			PublishedPostUUID: "",
			Category: &cateEntities.PostCategoryResDat{
				ID:   post.Category().ID,
				Name: post.Category().Name,
			},
			Tags: &cateEntities.PostTagResDat{
				ID:   post.Tags().ID,
				Tags: post.Tags().Tags,
			},
			Comments:  post.Comments(),
			Like:      post.Likes(),
			CreatedAt: post.CreatedAt.String(),
			UpdateAt:  post.UpdatedAt.String(),
		}
		if pbp, ok := post.PublishPostUUID(); ok {
			rp.PublishedPostUUID = pbp
		}

		pRes = append(pRes, rp)
	}
	return pRes
}
