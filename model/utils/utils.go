package utils

import (
	"net/http"

	"github.com/go-playground/validator"
	_errEntity "github.com/textures1245/BlogDuaaeeg-backend/error/entity"
)

func SchemaValidator[T any](req *T) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		return &_errEntity.CError{
			StatusCode: http.StatusBadRequest,
			Err:        errors,
		}
	}

	return nil
}

func Contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
