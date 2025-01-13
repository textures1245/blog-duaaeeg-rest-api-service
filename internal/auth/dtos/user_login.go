package dtos

import "github.com/golang-jwt/jwt/v4"

type UsersLoginRes struct {
	AccessToken string           `json:"access_token"`
	ExpiresAt   *jwt.NumericDate `json:"expires_at"`
	CreatedAt   string           `json:"created_at"`
	UpdatedAt   string           `json:"updated_at"`
}
