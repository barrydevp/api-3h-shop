package actions

import (
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
	"strings"
)

func GetListProductItem(queryProductItem *model.QueryProductItem) (*response.DataList, error) {
	queryString := ""
	var args []interface{}

	var where []string

	if queryProductItem.Id != nil {
		where = append(where, " _id=?")
		args = append(args, queryProductItem.Id)
	}
	if queryProductItem.ProductId != nil {
		if *queryProductItem.ProductId != 0 {
			if *queryProductItem.ProductId == -1 {
				where = append(where, " product_id IS NULL")
			} else {
				where = append(where, " product_id=?")
				args = append(args, queryProductItem.ProductId)
			}
		}
	}
	if queryProductItem.CreatedAtFrom != nil && queryProductItem.CreatedAtTo != nil {
		where = append(where, " created_at BETWEEN ? AND ?")
		args = append(args, queryProductItem.CreatedAtFrom, queryProductItem.CreatedAtTo)
	}
	if queryProductItem.UpdatedAtFrom != nil && queryProductItem.UpdatedAtTo != nil {
		where = append(where, " updated_at BETWEEN ? AND ?")
		args = append(args, queryProductItem.UpdatedAtFrom, queryProductItem.UpdatedAtTo)
	}
	if queryProductItem.ExpiredAtFrom != nil && queryProductItem.ExpiredAtTo != nil {
		where = append(where, " expired_at BETWEEN ? AND ?")
		args = append(args, queryProductItem.ExpiredAtFrom, queryProductItem.ExpiredAtTo)
	}

	if len(where) > 0 {
		queryString += "WHERE" + strings.Join(where, " AND") + "\n"
	}

	queryCountList := connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	}

	queryString += "ORDER BY _id ASC\n"

	queryProductItem.ParsePaging()
	queryString += "LIMIT ?, ?\n"

	args = append(args, queryProductItem.Offset)
	args = append(args, queryProductItem.Limit)

	queryList := connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	}

	resolveChan := make(chan interface{}, 2)
	rejectChan := make(chan error)

	go func() {
		data, err := factories.FindProductItem(&queryList)

		if err != nil {
			rejectChan <- err
		} else {
			resolveChan <- data
		}
	}()

	go func() {
		total, err := factories.CountProductItem(&queryCountList)

		if err != nil {
			rejectChan <- err
		}

		resolveChan <- total
	}()

	var data []*model.ProductItem
	var total int

	for i := 0; i < 2; i++ {
		select {
		case res := <-resolveChan:
			switch val := res.(type) {
			case []*model.ProductItem:
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
		Page:  *queryProductItem.Page,
		Limit: *queryProductItem.Limit,
	}, nil
}
