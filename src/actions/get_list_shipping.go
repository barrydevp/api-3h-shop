package actions

import (
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
	"strings"
)

func GetListShipping(queryShipping *model.QueryShipping) (*response.DataList, error) {
	queryString := ""
	var args []interface{}

	var where []string

	if queryShipping.Id != nil {
		where = append(where, " _id=?")
		args = append(args, queryShipping.Id)
	}
	if queryShipping.OrderId != nil {
		if *queryShipping.OrderId == -1 {
			where = append(where, " order_id IS NULL")
		} else {
			where = append(where, " order_id=?")
			args = append(args, queryShipping.OrderId)
		}
	}
	if queryShipping.Carrier != nil {
		where = append(where, " carrier LIKE ?")
		args = append(args, "%"+*queryShipping.Carrier+"%")
	}
	if queryShipping.Status != nil {
		where = append(where, " status=?")
		args = append(args, queryShipping.Status)
	}
	if queryShipping.CreatedAtFrom != nil && queryShipping.CreatedAtTo != nil {
		where = append(where, " created_at BETWEEN ? AND ?")
		args = append(args, queryShipping.CreatedAtFrom, queryShipping.CreatedAtTo)
	}
	if queryShipping.UpdatedAtFrom != nil && queryShipping.UpdatedAtTo != nil {
		where = append(where, " updated_at BETWEEN ? AND ?")
		args = append(args, queryShipping.UpdatedAtFrom, queryShipping.UpdatedAtTo)
	}
	if queryShipping.DeliveredAtFrom != nil && queryShipping.DeliveredAtTo != nil {
		where = append(where, " delivered_at BETWEEN ? AND ?")
		args = append(args, queryShipping.DeliveredAtFrom, queryShipping.DeliveredAtTo)
	}

	if len(where) > 0 {
		queryString += "WHERE" + strings.Join(where, " AND") + "\n"
	}

	queryCountList := connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	}

	queryString += "ORDER BY _id ASC\n"

	queryShipping.ParsePaging()
	queryString += "LIMIT ?, ?\n"

	args = append(args, queryShipping.Offset)
	args = append(args, queryShipping.Limit)

	queryList := connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	}

	resolveChan := make(chan interface{}, 2)
	rejectChan := make(chan error)

	go func() {
		data, err := factories.FindShipping(&queryList)

		if err != nil {
			rejectChan <- err
		} else {
			resolveChan <- data
		}
	}()

	go func() {
		total, err := factories.CountShipping(&queryCountList)

		if err != nil {
			rejectChan <- err
		}

		resolveChan <- total
	}()

	var data []*model.Shipping
	var total int

	for i := 0; i < 2; i++ {
		select {
		case res := <-resolveChan:
			switch val := res.(type) {
			case []*model.Shipping:
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
		Page:  *queryShipping.Page,
		Limit: *queryShipping.Limit,
	}, nil
}
