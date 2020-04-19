package actions

import (
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
	"strings"
)

func GetListCustomer(queryCustomer *model.QueryCustomer) (*response.DataList, error) {
	queryString := ""
	var args []interface{}

	var where []string

	if queryCustomer.Id != nil {
		where = append(where, " _id=?")
		args = append(args, queryCustomer.Id)
	}
	if queryCustomer.Phone != nil {
		where = append(where, " phone=?")
		args = append(args, queryCustomer.Phone)
	}
	if queryCustomer.Address != nil {
		where = append(where, " address LIKE ?")
		args = append(args, "%"+*queryCustomer.Address+"%")
	}
	if queryCustomer.FullName != nil {
		where = append(where, " full_name LIKE ?")
		args = append(args, "%"+*queryCustomer.FullName+"%")
	}
	if queryCustomer.Email != nil {
		where = append(where, " email LIKE ?")
		args = append(args, "%"+*queryCustomer.Email+"%")
	}
	if queryCustomer.UpdatedAtFrom != nil && queryCustomer.UpdatedAtTo != nil {
		where = append(where, " updated_at BETWEEN ? AND ?")
		args = append(args, queryCustomer.UpdatedAtFrom, queryCustomer.UpdatedAtTo)
	}

	if len(where) > 0 {
		queryString += "WHERE" + strings.Join(where, " AND") + "\n"
	}

	queryCountList := connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	}

	queryString += "ORDER BY _id ASC\n"

	queryCustomer.ParsePaging()
	queryString += "LIMIT ?, ?\n"

	args = append(args, queryCustomer.Offset)
	args = append(args, queryCustomer.Limit)

	queryList := connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	}

	resolveChan := make(chan interface{}, 2)
	rejectChan := make(chan error)

	go func() {
		data, err := factories.FindCustomer(&queryList)

		if err != nil {
			rejectChan <- err
		} else {
			resolveChan <- data
		}
	}()

	go func() {
		total, err := factories.CountCustomer(&queryCountList)

		if err != nil {
			rejectChan <- err
		}

		resolveChan <- total
	}()

	var data []*model.Customer
	var total int

	for i := 0; i < 2; i++ {
		select {
		case res := <-resolveChan:
			switch val := res.(type) {
			case []*model.Customer:
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
		Page:  *queryCustomer.Page,
		Limit: *queryCustomer.Limit,
	}, nil
}
