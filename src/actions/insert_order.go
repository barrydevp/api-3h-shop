package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
	"github.com/google/uuid"
	"strings"
)

func InsertOrder(body *model.BodyOrder) (*model.Order, error) {
	queryString := ""
	var args []interface{}

	var set []string

	_uuid, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	set = append(set, " session=?")
	args = append(args, _uuid.String())

	if body.Note != nil {
		set = append(set, " note=?")
		args = append(args, body.Note)
	}

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + "\n"
	} else {
		return nil, errors.New("invalid body")
	}

	id, err := factories.InsertOrder(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	if id == nil {
		return nil, errors.New("insert error")
	}

	return factories.FindOrderById(*id)
}
