package service

import (
	"fmt"
	"log"

	"github.com/textures1245/BlogDuaaeeg-backend/db"
	entityCate "github.com/textures1245/BlogDuaaeeg-backend/model/category/entity"
	postEntity "github.com/textures1245/BlogDuaaeeg-backend/model/post/entity"
	userEntity "github.com/textures1245/BlogDuaaeeg-backend/model/user/entity"
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
	if _, isNil := post.PublicationPost(); !isNil {
		if post.Published {
			err := u.PostRepo.UpdatePostToPublisher(post.UserUUID, post.UUID)
			if err != nil {
				fmt.Println(fmt.Errorf("%v", err))
				return nil, err
			}
		}
	}

	var pbpUuid = ""
	if uuid, ok := post.PublishPostUUID(); ok {
		pbpUuid = uuid
	}

	res := &postEntity.PostResDat{
		UUID:              post.UUID,
		UserUuid:          post.UserUUID,
		Title:             post.Title,
		Source:            post.Source,
		Published:         post.Published,
		SrcType:           string(post.SrcType),
		PublishedPostUUID: pbpUuid,
		Category:          cateResDat,
		Tags:              tagResDat,
	}

	return res, nil

}

func (u *postUse) OnUpdatePostAndTagByUUID(cateResDat *entityCate.PostCategoryResDat, uuid string, req *postEntity.PostReqDat) (*postEntity.PostResDat, error) {
	post, err := u.PostRepo.UpdatePostByUUID(cateResDat, uuid, req)
	if err != nil {
		return nil, err
	}

	// check if post marked as publish then link to publication post
	if _, isNil := post.PublicationPost(); !isNil {
		log.Print(post)
		if post.Published {
			err := u.PostRepo.UpdatePostToPublisher(post.UserUUID, post.UUID)
			log.Println("error on update post to publisher", err)
			if err != nil {
				return nil, err
			}
		}
	}

	tagUpdated, err := u.TagRepo.UpdateTags(cateResDat.ID, req.PostTag)
	if err != nil {
		return nil, err
	}

	pbpUuid, _, _ := prismaOptKeyRetrieve(post)

	res := &postEntity.PostResDat{
		UUID:              post.UUID,
		UserUuid:          post.UserUUID,
		Title:             post.Title,
		Source:            post.Source,
		Published:         post.Published,
		SrcType:           string(post.SrcType),
		PublishedPostUUID: pbpUuid,
		Category:          cateResDat,
		Tags: &entityCate.PostTagResDat{
			ID:   tagUpdated.ID,
			Tags: tagUpdated.Tags,
		},
	}
	return res, nil

}

func (u *postUse) OnFetchPostByUUID(uuid string) (*postEntity.PostResDat, error) {
	post, err := u.PostRepo.FetchPostByUUID(uuid)
	if err != nil {
		return nil, err
	}

	pbpUuid, cateM, tagM := prismaOptKeyRetrieve(post)

	res := &postEntity.PostResDat{
		UUID:              post.UUID,
		UserUuid:          post.UserUUID,
		Title:             post.Title,
		Source:            post.Source,
		Published:         post.Published,
		SrcType:           string(post.SrcType),
		PublishedPostUUID: pbpUuid,
		Category: &entityCate.PostCategoryResDat{
			ID:   cateM.ID,
			Name: cateM.Name,
		},
		Tags: &entityCate.PostTagResDat{
			ID:   tagM.ID,
			Tags: tagM.Tags,
		},
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

		_, cateM, tagM := prismaOptKeyRetrieve(post.Post())

		res = append(res, &postEntity.PostResDat{
			UUID:              post.UUID,
			UserUuid:          post.UserUUID,
			Title:             post.Post().Title,
			Source:            post.Post().Source,
			Published:         post.Post().Published,
			SrcType:           string(post.Post().SrcType),
			PublishedPostUUID: post.PostUUID,
			Category: &entityCate.PostCategoryResDat{
				ID:   cateM.ID,
				Name: cateM.Name,
			},
			Tags: &entityCate.PostTagResDat{
				ID:   tagM.ID,
				Tags: tagM.Tags,
			},
		})
	}
	return res, nil
}

// - this func should be called as private func
func (u *postUse) OnSubmitPostToPublisher(userUuid string, postUuid string) error {
	err := u.PostRepo.UpdatePostToPublisher(userUuid, postUuid)
	if err != nil {
		return err
	}
	return nil
}

func mapPostsDatToRes(pDat []db.PostModel, pRes []*postEntity.PostResDat) []*postEntity.PostResDat {
	for _, post := range pDat {

		pbpUuid, cateM, tagM := prismaOptKeyRetrieve(&post)

		pRes = append(pRes, &postEntity.PostResDat{
			UUID:              post.UUID,
			UserUuid:          post.UserUUID,
			Title:             post.Title,
			Source:            post.Source,
			Published:         post.Published,
			SrcType:           string(post.SrcType),
			PublishedPostUUID: pbpUuid,
			Category: &entityCate.PostCategoryResDat{
				ID: cateM.ID,
			},
			Tags: &entityCate.PostTagResDat{
				ID:   tagM.ID,
				Tags: tagM.Tags,
			},
		})
	}
	return pRes
}

func prismaOptKeyRetrieve(post *db.PostModel) (string, *db.PostCategoryModel, *db.PostTagModel) {

	var pbpUuid = ""
	if uuid, ok := post.PublishPostUUID(); ok {
		pbpUuid = uuid
	}
	var cateM *db.PostCategoryModel
	if cate, ok := post.Category(); ok {
		cateM = cate
	}
	var tagM *db.PostTagModel
	if tag, ok := post.Tags(); ok {
		tagM = tag
	}

	return pbpUuid, cateM, tagM

}
