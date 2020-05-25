package controllers

import (
	"errors"
	"strconv"

	"github.com/barrydev/api-3h-shop/src/actions"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
	"github.com/gin-gonic/gin"
)

func GetListUser(c *gin.Context) (interface{}, error) {
	var query model.QueryUser

	err := c.ShouldBindQuery(&query)

	if err != nil {
		return nil, err
	}

	return actions.GetListUser(&query)
}

func GetUserById(c *gin.Context) (interface{}, error) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

	if err != nil {
		return nil, err
	}

	return actions.GetUserById(userId)
}

func RegisterUser(c *gin.Context) (interface{}, error) {
	var body model.BodyUser

	err := c.ShouldBindJSON(&body)

	if err != nil {
		return nil, err
	}

	session, err := c.Cookie("3h.session")

	if err != nil {
		body.Session = nil
	} else {
		body.Session = &session
	}

	var userRole int64 = 11
	body.Role = &userRole

	return actions.InsertUser(&body)
}

func AuthenticateUser(c *gin.Context) (interface{}, error) {
	var body model.BodyUser

	err := c.ShouldBindJSON(&body)

	if err != nil {
		return nil, err
	}

	return actions.AuthenticateUser(&body)
}

func AuthenticateAdmin(c *gin.Context) (interface{}, error) {
	var body model.BodyUser

	err := c.ShouldBindJSON(&body)

	if err != nil {
		return nil, err
	}

	return actions.AuthenticateAdmin(&body)
}

func InsertUser(c *gin.Context) (interface{}, error) {
	var insertUser model.BodyUser

	err := c.ShouldBindJSON(&insertUser)

	if err != nil {
		return nil, err
	}

	return actions.InsertUser(&insertUser)
}

func UpdateProfile(c *gin.Context) (interface{}, error) {
	payload, ok := c.Get("payload_token")

	if !ok {
		return nil, errors.New("invalid payload_token")
	}

	payloadToken, ok := payload.(*factories.AccessTokenClaims)

	if !ok {
		return nil, errors.New("invalid payload_token")
	}

	var body model.BodyUser

	err := c.ShouldBindJSON(&body)

	if err != nil {
		return nil, err
	}

	return actions.UpdateUser(payloadToken.Id, &body)
}

func UpdateUser(c *gin.Context) (interface{}, error) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

	if err != nil {
		return nil, err
	}

	var body model.BodyUser

	err = c.ShouldBindJSON(&body)

	if err != nil {
		return nil, err
	}

	return actions.UpdateUser(userId, &body)
}

// func DeleteUser(c *gin.Context) (interface{}, error) {
// 	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return actions.DeleteUser(userId)
// }

func ChangeUserRole(c *gin.Context) (interface{}, error) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

	if err != nil {
		return nil, err
	}

	var body model.BodyUser

	err = c.ShouldBindJSON(&body)

	if err != nil {
		return nil, err
	}

	return actions.ChangeUserRole(userId, &body)
}

func ChangeUserPassword(c *gin.Context) (interface{}, error) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

	if err != nil {
		return nil, err
	}

	var body model.BodyUserChangePassword

	err = c.ShouldBindJSON(&body)

	if err != nil {
		return nil, err
	}

	return actions.ChangeUserPassword(userId, &body)
}

func ResetUserPassword(c *gin.Context) (interface{}, error) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

	if err != nil {
		return nil, err
	}

	var body model.BodyUser

	err = c.ShouldBindJSON(&body)

	if err != nil {
		return nil, err
	}

	return actions.ResetUserPassword(userId, &body)
}
