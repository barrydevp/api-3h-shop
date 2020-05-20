package factories

import (
	"errors"
	"fmt"

	"github.com/barrydev/api-3h-shop/src/constants"
	"github.com/barrydev/api-3h-shop/src/helpers"
	"github.com/barrydev/api-3h-shop/src/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

func RetriveAccessTokenPayload(c *gin.Context) (*AccessTokenClaims, error) {
	tokenString := helpers.GetAccessToken(c)

	if tokenString == "" {
		return nil, errors.New("invalid token")
	}

	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(constants.SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AccessTokenClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
