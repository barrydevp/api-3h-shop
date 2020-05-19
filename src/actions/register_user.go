package actions

import (
	"github.com/barrydev/api-3h-shop/src/model"
)

func RegisterUser(body *model.BodyUser) (*model.User, error) {
	var userRole int64 = 11
	body.Role = &userRole

	return InsertUser(body)
}
