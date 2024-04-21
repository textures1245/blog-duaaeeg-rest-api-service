package handler

import (
	"errors"

	"github.com/textures1245/BlogDuaaeeg-backend/db"
	"github.com/textures1245/BlogDuaaeeg-backend/error/entity"
)

type HandleUse struct {
	HandleRepo entity.ErrorHandler
}

func NewHandler(errHandle entity.ErrorHandler) entity.ErrorHandler {
	return &HandleUse{
		HandleRepo: errHandle,
	}
}

func (h *HandleUse) PrismaAuthHandle(err entity.CError) *entity.CError {

	if errors.Is(err.Err, db.ErrNotFound) {
		err.Err = errors.New("UserNotFound")
		return &err
	} else if _, uniqueErr := db.IsErrUniqueConstraint(err.Err); uniqueErr {
		err.Err = errors.New("EmailAlreadyExists")
		return &err
	} else {
		return &err
	}
}
