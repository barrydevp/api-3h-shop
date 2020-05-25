package actions

import (
	"errors"
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func ChangeUserRole(userId int64, body *model.BodyUser) (bool, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.Role != nil {
		if !model.IsSupportedRole(*body.Role) {
			return false, errors.New("we don't support that role")
		}
		set = append(set, " role=?")
		args = append(args, body.Role)
	} else {
		return false, errors.New("user's role is required")
	}

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + "\n"
	} else {
		return false, errors.New("nothing has change")
	}

	user, err := factories.FindUserById(userId)

	if err != nil {
		return false, err
	}

	if user == nil {
		return false, errors.New("user does not exists")
	}

	queryString += "WHERE _id=?"
	args = append(args, userId)

	rowEffected, err := factories.UpdateUser(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return false, err
	}

	if rowEffected == nil {
		return false, errors.New("nothing has change")
	}

	return true, nil
}
