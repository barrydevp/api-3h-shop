package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
)

func DeleteOrderItem(orderItemId int64) (bool, error) {
	rowEffected, err := factories.DeleteOrderItem(&connect.QueryMySQL{
		QueryString: "WHERE _id=?",
		Args:        []interface{}{&orderItemId},
	})

	if err != nil {
		return false, err
	}

	if rowEffected == nil || *rowEffected == 0 {
		return false, errors.New("update error")
	}

	return true, nil
}
