package actions

import (
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetListCoupon(queryCoupon *model.QueryCoupon) (*response.DataList, error) {
	queryString := ""
	var args []interface{}

	var where []string

	if queryCoupon.Id != nil {
		where = append(where, " _id=?")
		args = append(args, queryCoupon.Id)
	}
	if queryCoupon.Code != nil {
		where = append(where, " code=?")
		args = append(args, queryCoupon.Code)
	}
	if queryCoupon.Discount != nil {
		where = append(where, " discount=?")
		args = append(args, queryCoupon.Discount)
	}
	if queryCoupon.UpdatedAtFrom != nil && queryCoupon.UpdatedAtTo != nil {
		where = append(where, " updated_at BETWEEN ? AND ?")
		args = append(args, queryCoupon.UpdatedAtFrom, queryCoupon.UpdatedAtTo)
	}
	if queryCoupon.ExpiresAtFrom != nil && queryCoupon.ExpiresAtTo != nil {
		where = append(where, " expires_at BETWEEN ? AND ?")
		args = append(args, queryCoupon.ExpiresAtFrom, queryCoupon.ExpiresAtTo)
	}

	if len(where) > 0 {
		queryString += "WHERE" + strings.Join(where, " AND") + "\n"
	}

	queryCountList := connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	}

	queryCoupon.ParseSort()
	queryString += "ORDER BY " + *queryCoupon.OrderBy + "\n"

	queryCoupon.ParsePaging()
	queryString += "LIMIT ?, ?\n"

	args = append(args, queryCoupon.Offset)
	args = append(args, queryCoupon.Limit)

	queryList := connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	}

	resolveChan := make(chan interface{}, 2)
	rejectChan := make(chan error)

	go func() {
		data, err := factories.FindCoupon(&queryList)

		if err != nil {
			rejectChan <- err
		} else {
			resolveChan <- data
		}
	}()

	go func() {
		total, err := factories.CountCoupon(&queryCountList)

		if err != nil {
			rejectChan <- err
		}

		resolveChan <- total
	}()

	var data []*model.Coupon
	var total int

	for i := 0; i < 2; i++ {
		select {
		case res := <-resolveChan:
			switch val := res.(type) {
			case []*model.Coupon:
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
		Page:  *queryCoupon.Page,
		Limit: *queryCoupon.Limit,
	}, nil
}
