package actions

import (
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
	"strings"
)

func GetListOrderItem(queryOrderItem *model.QueryOrderItem) (*response.DataList, error) {
	queryString := ""
	var args []interface{}

	var where []string

	if queryOrderItem.Id != nil {
		where = append(where, " _id=?")
		args = append(args, queryOrderItem.Id)
	}
	if queryOrderItem.OrderId != nil {
		if *queryOrderItem.OrderId == 0 {
			where = append(where, " order_id IS NULL")
		} else {
			where = append(where, " order_id=?")
			args = append(args, queryOrderItem.OrderId)
		}
	}
	if queryOrderItem.ProductId != nil {
		if *queryOrderItem.ProductId == 0 {
			where = append(where, " product_id IS NULL")
		} else {
			where = append(where, " product_id=?")
			args = append(args, queryOrderItem.ProductId)
		}
	}
	if queryOrderItem.ProductItemId != nil {
		if *queryOrderItem.ProductItemId == 0 {
			where = append(where, " product_item_id IS NULL")
		} else {
			where = append(where, " product_item_id=?")
			args = append(args, queryOrderItem.ProductItemId)
		}
	}
	if queryOrderItem.CreatedAtFrom != nil && queryOrderItem.CreatedAtTo != nil {
		where = append(where, " created_at BETWEEN ? AND ?")
		args = append(args, queryOrderItem.CreatedAtFrom, queryOrderItem.CreatedAtTo)
	}
	if queryOrderItem.UpdatedAtFrom != nil && queryOrderItem.UpdatedAtTo != nil {
		where = append(where, " updated_at BETWEEN ? AND ?")
		args = append(args, queryOrderItem.UpdatedAtFrom, queryOrderItem.UpdatedAtTo)
	}

	if len(where) > 0 {
		queryString += "WHERE" + strings.Join(where, " AND") + "\n"
	}

	queryCountList := connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	}

	queryString += "ORDER BY _id ASC\n"

	queryOrderItem.ParsePaging()
	queryString += "LIMIT ?, ?\n"

	args = append(args, queryOrderItem.Offset)
	args = append(args, queryOrderItem.Limit)

	queryList := connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	}

	resolveChan := make(chan interface{}, 2)
	rejectChan := make(chan error)

	go func() {
		data, err := factories.FindOrderItem(&queryList)

		if err != nil {
			rejectChan <- err
		} else {
			resolveChan <- data
		}
	}()

	go func() {
		total, err := factories.CountOrderItem(&queryCountList)

		if err != nil {
			rejectChan <- err
		}

		resolveChan <- total
	}()

	var data []*model.OrderItem
	var total int

	for i := 0; i < 2; i++ {
		select {
		case res := <-resolveChan:
			switch val := res.(type) {
			case []*model.OrderItem:
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
		Page:  *queryOrderItem.Page,
		Limit: *queryOrderItem.Limit,
	}, nil
}
