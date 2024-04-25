package service

import (
	"fmt"

	"github.com/textures1245/BlogDuaaeeg-backend/api/db"
	entityCate "github.com/textures1245/BlogDuaaeeg-backend/api/model/category/entity"
	postEntity "github.com/textures1245/BlogDuaaeeg-backend/api/model/post/entity"
	userEntity "github.com/textures1245/BlogDuaaeeg-backend/api/model/user/entity"
)

type postUse struct {
	PostRepo  postEntity.PostRepository
	UsersRepo userEntity.UsersRepository
	TagRepo   entityCate.PostTagRepository
}

func NewPostService(postRepo postEntity.PostRepository, usersRepo userEntity.UsersRepository, tagRepo entityCate.PostTagRepository) postEntity.PostService {
	return &postUse{
		PostRepo:  postRepo,
		UsersRepo: usersRepo,
		TagRepo:   tagRepo,
	}
}

func (u *postUse) OnCreateNewPost(cateResDat *entityCate.PostCategoryResDat, tagResDat *entityCate.PostTagResDat, req *postEntity.PostReqDat) (*postEntity.PostResDat, error) {
	post, err := u.PostRepo.CreatePost(cateResDat, tagResDat, req)
	if err != nil {
		return nil, err
	}

	// check if post marked as publish then link to publication post
	pbpUuid := ""
	if oldUuid, isNil := post.PublishPostUUID(); !isNil {
		if post.Published {
			uuid, err := u.PostRepo.UpdatePostToPublisher(post.UserUUID, post.UUID)
			if err != nil {
				fmt.Println(fmt.Errorf("%v", err))
				return nil, err
			}
			pbpUuid = uuid
		}
	} else {
		pbpUuid = oldUuid
	}

	res := &postEntity.PostResDat{
		UUID:              post.UUID,
		UserUuid:          post.UserUUID,
		Title:             post.Title,
		Source:            post.Source,
		Published:         post.Published,
		SrcType:           string(post.SrcType),
		PublishedPostUUID: pbpUuid,
		Category: &entityCate.PostCategoryResDat{
			ID:   post.Category().ID,
			Name: post.Category().Name,
		},
		Tags: &entityCate.PostTagResDat{
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

func (u *postUse) OnUpdatePostAndTagByUUID(cateResDat *entityCate.PostCategoryResDat, uuid string, req *postEntity.PostReqDat) (*postEntity.PostResDat, error) {
	post, err := u.PostRepo.UpdatePostByUUID(cateResDat, uuid, req)
	if err != nil {
		return nil, err
	}

	// check if post marked as publish then link to publication post
	pbpUuid := ""
	if oldUuid, isNil := post.PublishPostUUID(); !isNil {
		if post.Published {
			uuid, err := u.PostRepo.UpdatePostToPublisher(post.UserUUID, post.UUID)
			if err != nil {
				fmt.Println(fmt.Errorf("%v", err))
				return nil, err
			}
			pbpUuid = uuid
		}
	} else {
		pbpUuid = oldUuid
	}

	tagUpdated, err := u.TagRepo.UpdateTags(cateResDat.ID, req.PostTag)
	if err != nil {
		return nil, err
	}

	res := &postEntity.PostResDat{
		UUID:              post.UUID,
		UserUuid:          post.UserUUID,
		Title:             post.Title,
		Source:            post.Source,
		Published:         post.Published,
		SrcType:           string(post.SrcType),
		PublishedPostUUID: pbpUuid,
		Category: &entityCate.PostCategoryResDat{
			ID:   post.Category().ID,
			Name: post.Category().Name,
		},
		Tags: &entityCate.PostTagResDat{
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

func (u *postUse) OnFetchPostByUUID(uuid string) (*postEntity.PostResDat, error) {
	post, err := u.PostRepo.FetchPostByUUID(uuid)
	if err != nil {
		return nil, err
	}

	// pbpUuid, cateM, tagM := prismaOptKeyRetrieve(post)

	res := &postEntity.PostResDat{
		UUID:              post.UUID,
		UserUuid:          post.UserUUID,
		Title:             post.Title,
		Source:            post.Source,
		Published:         post.Published,
		SrcType:           string(post.SrcType),
		PublishedPostUUID: "",
		Category: &entityCate.PostCategoryResDat{
			ID:   post.Category().ID,
			Name: post.Category().Name,
		},
		Tags: &entityCate.PostTagResDat{
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

func (u *postUse) OnFetchOwnerPosts(userUuid string) ([]*postEntity.PostResDat, error) {
	posts, err := u.PostRepo.FetchPostByUserUUID(userUuid)
	if err != nil {
		return nil, err
	}

	var res []*postEntity.PostResDat
	res = mapPostsDatToRes(posts, res)
	return res, nil

}

func (u *postUse) OnFetchPublisherPosts(opts *postEntity.FetchPostOptReq) ([]*postEntity.PostResDat, error) {
	posts, err := u.PostRepo.FetchPublisherPosts(opts)
	if err != nil {
		return nil, err
	}

	var res []*postEntity.PostResDat
	for _, post := range posts {

		res = append(res, &postEntity.PostResDat{
			UUID:      post.UUID,
			UserUuid:  post.UserUUID,
			Title:     post.Post().Title,
			Source:    post.Post().Source,
			Published: post.Post().Published,
			SrcType:   string(post.Post().SrcType),
			PostUUID:  post.PostUUID,
			Category: &entityCate.PostCategoryResDat{
				ID:   post.Post().Category().ID,
				Name: post.Post().Category().Name,
			},
			Tags: &entityCate.PostTagResDat{
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
		return "", err
	}
	return pbpUuid, nil
}

func (u *postUse) OnDeletePostByUUID(postUuid string) error {

	err := u.PostRepo.DeletePostByUUID(postUuid)
	if err != nil {
		return err
	}

	return nil

}

func mapPostsDatToRes(pDat []db.PostModel, pRes []*postEntity.PostResDat) []*postEntity.PostResDat {
	for _, post := range pDat {

		rp := &postEntity.PostResDat{
			UUID:              post.UUID,
			UserUuid:          post.UserUUID,
			Title:             post.Title,
			Source:            post.Source,
			Published:         post.Published,
			SrcType:           string(post.SrcType),
			PublishedPostUUID: "",
			Category: &entityCate.PostCategoryResDat{
				ID:   post.Category().ID,
				Name: post.Category().Name,
			},
			Tags: &entityCate.PostTagResDat{
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
