package helpers

import (
	"fmt"
	"strings"

	"github.com/barrydev/api-3h-shop/src/constants"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateJwtToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SECRET_KEY))
}

func GetAccessToken(c *gin.Context) string {
	bearToken := c.GetHeader("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	tokenFromSession, err := c.Cookie("access.token")

	if err != nil {
		return ""
	}

	return tokenFromSession
}

func GetRefreshToken(c *gin.Context) string {
	refreshToken, err := c.Cookie("refresh.token")

	if err != nil {
		return ""
	}

	return refreshToken
}

func RetriveJwtToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := GetAccessToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(constants.SECRET_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func IsValidJwtToken(c *gin.Context) error {
	token, err := RetriveJwtToken(c)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}
