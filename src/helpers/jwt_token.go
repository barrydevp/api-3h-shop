package helpers

import (
	"github.com/dgrijalva/jwt-go"
)

type TokenPayload struct {
	Id   int64 `json:"_id"`
	Role int64 `json:"role"`
	jwt.StandardClaims
}

func GenerateJwtToken(claims interface{}) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString()
}
