package actions

import (
	"errors"

	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetUserById(userId int64) (*model.User, error) {
	user, err := factories.FindUserById(userId)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user does not exists")
	}

	return user, nil
}
