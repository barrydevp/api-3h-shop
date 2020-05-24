package actions

import (
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetListOrder(queryOrder *model.QueryOrder) (*response.DataList, error) {
	queryString := ""
	var args []interface{}

	var where []string

	if queryOrder.Id != nil {
		where = append(where, " _id=?")
		args = append(args, queryOrder.Id)
	}
	if queryOrder.CustomerId != nil {
		if *queryOrder.CustomerId != 0 {
			if *queryOrder.CustomerId == -1 {
				where = append(where, " order_id IS NULL")
			} else {
				where = append(where, " order_id=?")
				args = append(args, queryOrder.CustomerId)
			}
		}
	}
	if queryOrder.PaymentStatus != nil {
		where = append(where, " payment_status=?")
		args = append(args, *queryOrder.PaymentStatus)
	}
	if queryOrder.FulfillmentStatus != nil {
		where = append(where, " fulfillment_status=?")
		args = append(args, *queryOrder.FulfillmentStatus)
	}
	if queryOrder.Note != nil {
		where = append(where, " note LIKE ?")
		args = append(args, "%"+*queryOrder.Note+"%")
	}
	if queryOrder.CreatedAtFrom != nil && queryOrder.CreatedAtTo != nil {
		where = append(where, " created_at BETWEEN ? AND ?")
		args = append(args, queryOrder.CreatedAtFrom, queryOrder.CreatedAtTo)
	}
	if queryOrder.UpdatedAtFrom != nil && queryOrder.UpdatedAtTo != nil {
		where = append(where, " updated_at BETWEEN ? AND ?")
		args = append(args, queryOrder.UpdatedAtFrom, queryOrder.UpdatedAtTo)
	}
	if queryOrder.PaidAtFrom != nil && queryOrder.PaidAtTo != nil {
		where = append(where, " paid_at BETWEEN ? AND ?")
		args = append(args, queryOrder.PaidAtFrom, queryOrder.PaidAtFrom)
	}
	if queryOrder.FulfilledAtFrom != nil && queryOrder.FulfilledAtTo != nil {
		where = append(where, " fulfillment_at BETWEEN ? AND ?")
		args = append(args, queryOrder.FulfilledAtFrom, queryOrder.FulfilledAtTo)
	}
	if queryOrder.CancelledAtFrom != nil && queryOrder.CancelledAtTo != nil {
		where = append(where, " cancelled_at BETWEEN ? AND ?")
		args = append(args, queryOrder.CancelledAtFrom, queryOrder.CancelledAtTo)
	}

	if len(where) > 0 {
		queryString += "WHERE" + strings.Join(where, " AND") + "\n"
	}

	queryCountList := connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	}

	queryOrder.ParseSort()
	queryString += "ORDER BY " + *queryOrder.OrderBy + "\n"

	queryOrder.ParsePaging()
	queryString += "LIMIT ?, ?\n"

	args = append(args, queryOrder.Offset)
	args = append(args, queryOrder.Limit)

	queryList := connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	}

	resolveChan := make(chan interface{}, 2)
	rejectChan := make(chan error)

	go func() {
		data, err := factories.FindOrder(&queryList)

		if err != nil {
			rejectChan <- err
		} else {
			resolveChan <- data
		}
	}()

	go func() {
		total, err := factories.CountOrder(&queryCountList)

		if err != nil {
			rejectChan <- err
		}

		resolveChan <- total
	}()

	var data []*model.Order
	var total int

	for i := 0; i < 2; i++ {
		select {
		case res := <-resolveChan:
			switch val := res.(type) {
			case []*model.Order:
				data = val
			case int:
				total = val
			}

		case err := <-rejectChan:
			return nil, err
		}
	}

	return &response.DataList{
		Data:  data,
		Total: total,
		Page:  *queryOrder.Page,
		Limit: *queryOrder.Limit,
	}, nil
}
