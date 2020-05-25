package actions

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func ChangeUserPassword(userId int64, body *model.BodyUserChangePassword) (bool, error) {
	queryString := ""
	var args []interface{}

	var set []string

	foundUser, err := factories.FindUserById(userId)

	if err != nil {
		return false, err
	}

	if foundUser == nil {
		return false, errors.New("user does not exists")
	}

	if body.OldPassword != nil {
		hashPassword := md5.Sum([]byte(*body.OldPassword))
		if hex.EncodeToString(hashPassword[:]) != *foundUser.RawPassword {
			return false, errors.New("wrong password")
		}
	} else {
		return false, errors.New("old_password is required")
	}

	if body.NewPassword != nil {
		if *body.NewPassword == *body.OldPassword {
			return false, errors.New("nothing has change")
		}

		hashPassword := md5.Sum([]byte(*body.NewPassword))
		set = append(set, " password=?")
		args = append(args, hex.EncodeToString(hashPassword[:]))
	} else {
		return false, errors.New("new_password is required")
	}

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + "\n"
	} else {
		return false, errors.New("nothing has change")
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
		return false, errors.New("update error")
	}

	return true, nil
}
