package factories

import (
	"github.com/barrydev/api-3h-shop/src/model"
	"github.com/dgrijalva/jwt-go"
)

type AccessTokenClaims struct {
	Id   int64 `json:"_id"`
	Role int64 `json:"role"`
	jwt.StandardClaims
}

type ResponseAccessToken struct {
	User        *model.User `json:"user"`
	AccessToken *string     `json:"access_token"`
}
