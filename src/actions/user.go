package actions

import (
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetListUser() ([]*model.User, error) {
	return factories.GetListUser()
}
