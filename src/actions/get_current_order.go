package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetCurrentOrder(current *model.Current) (*model.CurrentResponse, error) {
	var order *model.Order

	if current.Session != nil {
		_order, err := factories.FindOneOrder(
			&connect.QueryMySQL{
				QueryString: "WHERE session=? AND status='pending'",
				Args:        []interface{}{current.Session},
			})

		if err != nil {
			return nil, err
		}

		order = _order
	}

	if order != nil {
		return &model.CurrentResponse{
			Order:     order,
			Session:   current.Session,
		}, nil
	}

	session, err := factories.GenerateNewSessionToken()

	if err != nil {
		return nil, err
	}

	current.Session = &session

	id, err := factories.InsertOrder(&connect.QueryMySQL{
		QueryString: "SET session=?, created_at=NOW() \n",
		Args:        []interface{}{&session},
	})

	if err != nil {
		return nil, err
	}

	order, err = factories.FindOrderById(*id)

	if order == nil {
		return nil, errors.New("special error")
	}

	return &model.CurrentResponse{
		Order:     order,
		Session:   current.Session,
	}, nil
}
