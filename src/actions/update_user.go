package actions

import (
	"errors"
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func UpdateUser(userId int64, body *model.BodyUser) (*model.User, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.Name != nil {
		set = append(set, " name=?")
		args = append(args, body.Name)
	}

	// if body.Phone != nil {
	// 	set = append(set, " phone=?")
	// 	args = append(args, body.Phone)
	// }

	// if body.Password != nil {
	// 	hashPassword := md5.Sum([]byte(*body.Password))
	// 	set = append(set, " password=?")
	// 	args = append(args, string(hashPassword[:]))
	// }

	if body.Address != nil {
		set = append(set, " address=?")
		args = append(args, body.Address)
	}

	// if body.Session != nil {
	// 	set = append(set, " session=?")
	// 	args = append(args, body.Session)
	// }

	// if body.Role != nil {
	// 	set = append(set, " role=?")
	// 	args = append(args, body.Role)
	// }

	// if body.Status != nil {
	// 	set = append(set, " status=?")
	// 	args = append(args, body.Status)
	// }

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + "\n"
	} else {
		user, err := factories.FindUserById(userId)

		if err != nil {
			return nil, err
		}

		if user == nil {
			return nil, errors.New("user does not exists")
		}

		return user, nil
	}

	queryString += "WHERE _id=?"
	args = append(args, userId)

	rowEffected, err := factories.UpdateUser(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	if rowEffected == nil {
		return nil, errors.New("update error")
	}

	return GetUserById(userId)
}
