package actions

import (
	"errors"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func NewSession(current *model.Current) (*model.CurrentResponse, error) {
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

	order, err := factories.FindOrderById(*id)
	if order == nil {
		return nil, errors.New("special error")
	}

	if err != nil {
		return nil, err
	}

	return &model.CurrentResponse{
		Order:     order,
		Session:   &session,
		TotalItem: 0,
	}, nil
}

func _getCurrentOrderAuth(current *model.Current, payloadToken *factories.AccessTokenClaims, needUpdateUser *bool) (*model.CurrentResponse, error) {
	var err error

	user, err := factories.FindOneUser(
		&connect.QueryMySQL{
			QueryString: "WHERE _id=?",
			Args:        []interface{}{payloadToken.Id},
		})

	if err != nil {
		return nil, err
	}

	if user.Session == nil && current.Session == nil {
		return NewSession(current)
	}

	var mapTotal map[string]int
	var mapOrders map[string]*model.Order

	if user.Session == nil {
		mapTotal, mapOrders, err = factories.GetOrderAndTotalItemWithSession(
			&connect.QueryMySQL{
				QueryString: "WHERE session=? AND status='pending'",
				Args:        []interface{}{current.Session},
			})

		if err != nil {
			return nil, err
		}

		total, ok := mapTotal[*current.Session]

		if !ok {
			return NewSession(current)
		}

		order := mapOrders[*current.Session]

		return &model.CurrentResponse{
			Order:     order,
			Session:   current.Session,
			TotalItem: total,
		}, nil
	}

	if current.Session == nil {
		mapTotal, mapOrders, err = factories.GetOrderAndTotalItemWithSession(
			&connect.QueryMySQL{
				QueryString: "WHERE session=? AND status='pending'",
				Args:        []interface{}{current.Session},
			})

		if err != nil {
			return nil, err
		}

		*needUpdateUser = false
		total, ok := mapTotal[*user.Session]

		if !ok {
			return NewSession(current)
		}

		order := mapOrders[*user.Session]

		return &model.CurrentResponse{
			Order:     order,
			Session:   user.Session,
			TotalItem: total,
		}, nil
	}

	mapTotal, mapOrders, err = factories.GetOrderAndTotalItemWithSession(
		&connect.QueryMySQL{
			QueryString: "WHERE session IN(?, ?) AND status='pending'",
			Args:        []interface{}{current.Session, user.Session},
		})

	if err != nil {
		return nil, err
	}

	totalItemOfUser, okUser := mapTotal[*user.Session]
	totalItemOfCurrent, okCurrent := mapTotal[*user.Session]

	if !okUser && !okCurrent {
		return NewSession(current)
	}

	if okUser {
		*needUpdateUser = false
		order := mapOrders[*user.Session]

		return &model.CurrentResponse{
			Order:     order,
			Session:   user.Session,
			TotalItem: totalItemOfUser,
		}, nil
	}

	order := mapOrders[*current.Session]

	return &model.CurrentResponse{
		Order:     order,
		Session:   current.Session,
		TotalItem: totalItemOfCurrent,
	}, nil
}

func GetCurrentOrderAuth(current *model.Current, payloadToken *factories.AccessTokenClaims) (*model.CurrentResponse, error) {
	needUpdateUser := true
	response, err := _getCurrentOrderAuth(current, payloadToken, &needUpdateUser)

	if err != nil {
		return nil, err
	}

	if needUpdateUser {
		_, err = factories.UpdateUser(
			&connect.QueryMySQL{
				QueryString: "SET session=? \n WHERE _id=?",
				Args:        []interface{}{response.Session, payloadToken.Id},
			})

		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

func GetCurrentOrderV2(current *model.Current) (*model.CurrentResponse, error) {
	if current.Session != nil {
		mapTotal, mapOrders, err := factories.GetOrderAndTotalItemWithSession(
			&connect.QueryMySQL{
				QueryString: "WHERE session=? AND status='pending'",
				Args:        []interface{}{current.Session},
			})

		if err != nil {
			return nil, err
		}

		total, ok := mapTotal[*current.Session]

		if !ok {
			return NewSession(current)
		}

		order := mapOrders[*current.Session]

		return &model.CurrentResponse{
			Order:     order,
			Session:   current.Session,
			TotalItem: total,
		}, nil
	}

	return NewSession(current)
}

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
		totalItem, err := factories.CountOrderItem(
			&connect.QueryMySQL{
				QueryString: "WHERE order_id=?",
				Args:        []interface{}{order.Id},
			})

		if err != nil {
			return nil, err
		}

		return &model.CurrentResponse{
			Order:     order,
			Session:   current.Session,
			TotalItem: totalItem,
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
		TotalItem: 0,
	}, nil
}
