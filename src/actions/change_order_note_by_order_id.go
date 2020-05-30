package actions

import (
	"errors"
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func ChangeOrderNoteByOrderId(orderId int64, body *model.BodyOrder) (bool, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.Note != nil {
		set = append(set, " note=?")
		args = append(args, body.Note)
	} else {
		return false, errors.New("require order's note")
	}

	set = append(set, " status=change_note")

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + "\n"
	}

	queryString += "WHERE _id=?"
	args = append(args, orderId)

	rowEffected, err := factories.UpdateOrder(&connect.QueryMySQL{
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
