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

func InsertUser(body *model.BodyUser) (*model.User, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.Email == nil {
		return nil, errors.New("user's email is required")
	} else {
		set = append(set, " email=?")
		args = append(args, body.Email)
	}

	if body.Phone != nil {
		set = append(set, " phone=?")
		args = append(args, body.Phone)
	} else {
		return nil, errors.New("user's phone is required")
	}

	user, err := factories.FindOneUser(&connect.QueryMySQL{
		QueryString: "WHERE email=? OR phone=?",
		Args:        []interface{}{body.Email, body.Phone},
	})

	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("email or phone is exists")
	}

	if body.Password == nil {
		return nil, errors.New("user's password is required")
	} else {
		hashPassword := md5.Sum([]byte(*body.Password))
		set = append(set, " password=?")
		args = append(args, hex.EncodeToString(hashPassword[:]))
	}

	if body.Name != nil {
		set = append(set, " name=?")
		args = append(args, body.Name)
	} else {
		return nil, errors.New("user's name is required")
	}

	if body.Address != nil {
		set = append(set, " address=?")
		args = append(args, body.Address)
	} else {
		return nil, errors.New("user's address is required")
	}

	if body.Session != nil {
		set = append(set, " session=?")
		args = append(args, body.Session)
	}

	if body.Role != nil {
		set = append(set, " role=?")
		args = append(args, body.Role)
	} else {
		return nil, errors.New("user's role is required")
	}

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + ", created_at=NOW() \n"

	} else {
		return nil, errors.New("invalid body")
	}

	id, err := factories.InsertUser(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	if id == nil {
		return nil, errors.New("insert error")
	}

	return factories.FindUserById(*id)
}
