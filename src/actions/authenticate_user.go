package actions

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/helpers"
	"github.com/barrydev/api-3h-shop/src/model"
	"github.com/dgrijalva/jwt-go"
)

func AuthenticateUser(body *model.BodyUser) (*factories.ResponseAccessToken, error) {
	queryString := ""
	var args []interface{}

	var where []string

	if body.Email == nil {
		return nil, errors.New("user's email is required")
	} else {
		where = append(where, " email=?")
		args = append(args, body.Email)
	}

	if body.Password == nil {
		return nil, errors.New("user's password is required")
	}

	if len(where) > 0 {
		queryString += "WHERE" + strings.Join(where, " AND") + "\n"
	} else {
		return nil, errors.New("invalid body")
	}

	foundUser, err := factories.FindOneUser(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	if foundUser == nil {
		return nil, errors.New("user not found")
	}

	hashPassword := md5.Sum([]byte(*body.Password))

	if hex.EncodeToString(hashPassword[:]) != *foundUser.RawPassword {
		return nil, errors.New("wrong password")
	}

	claims := factories.AccessTokenClaims{
		*foundUser.Id,
		*foundUser.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}

	accessToken, err := helpers.GenerateJwtToken(claims)

	if err != nil {
		return nil, err
	}

	return &factories.ResponseAccessToken{
		User:        foundUser,
		AccessToken: &accessToken,
	}, nil
}
