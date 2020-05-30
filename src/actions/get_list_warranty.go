package actions

import (
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetListWarranty(queryWarranty *model.QueryWarranty) (*response.DataList, error) {
	queryString := ""
	var args []interface{}

	var where []string

	if queryWarranty.Id != nil {
		where = append(where, " _id=?")
		args = append(args, queryWarranty.Id)
	}
	if queryWarranty.Code != nil {
		where = append(where, " code=?")
		args = append(args, queryWarranty.Code)
	}
	if queryWarranty.Month != nil {
		where = append(where, " month=?")
		args = append(args, queryWarranty.Month)
	}
	if queryWarranty.Trial != nil {
		where = append(where, " trial=?")
		args = append(args, queryWarranty.Trial)
	}
	if queryWarranty.Status != nil {
		where = append(where, " status=?")
		args = append(args, queryWarranty.Status)
	}
	if queryWarranty.CategoryId != nil {
		where = append(where, " (category_id=? OR EXISTS(SELECT _id FROM categories WHERE _id=? AND parent_id=warranties.category_id))")
		args = append(args, queryWarranty.CategoryId, queryWarranty.CategoryId)
	}

	if len(where) > 0 {
		queryString += "WHERE" + strings.Join(where, " AND") + "\n"
	}

	queryCountList := connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	}

	queryWarranty.ParseSort()
	queryString += "ORDER BY " + *queryWarranty.OrderBy + "\n"

	queryWarranty.ParsePaging()
	queryString += "LIMIT ?, ?\n"

	args = append(args, queryWarranty.Offset)
	args = append(args, queryWarranty.Limit)

	queryList := connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	}

	resolveChan := make(chan interface{}, 2)
	rejectChan := make(chan error)

	go func() {
		data, err := factories.FindWarranty(&queryList)

		if err != nil {
			rejectChan <- err
		} else {
			resolveChan <- data
		}
	}()

	go func() {
		total, err := factories.CountWarranty(&queryCountList)

		if err != nil {
			rejectChan <- err
		}

		resolveChan <- total
	}()

	var data []*model.Warranty
	var total int

	for i := 0; i < 2; i++ {
		select {
		case res := <-resolveChan:
			switch val := res.(type) {
			case []*model.Warranty:
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
		Page:  *queryWarranty.Page,
		Limit: *queryWarranty.Limit,
	}, nil
}
