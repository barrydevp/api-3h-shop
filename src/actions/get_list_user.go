package actions

import (
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetListUser(queryUser *model.QueryUser) (*response.DataList, error) {
	queryString := ""
	var args []interface{}

	var where []string

	if queryUser.Id != nil {
		where = append(where, " _id=?")
		args = append(args, queryUser.Id)
	}
	if queryUser.Email != nil {
		where = append(where, " email LIKE ?")
		args = append(args, "%"+*queryUser.Email+"%")
	}
	if queryUser.Name != nil {
		where = append(where, " name LIKE ?")
		args = append(args, "%"+*queryUser.Name+"%")
	}
	if queryUser.Address != nil {
		where = append(where, " address LIKE ?")
		args = append(args, "%"+*queryUser.Address+"%")
	}
	if queryUser.Phone != nil {
		where = append(where, " phone LIKE ?")
		args = append(args, "%"+*queryUser.Phone+"%")
	}
	if queryUser.Status != nil {
		where = append(where, " status=?")
		args = append(args, queryUser.Status)
	}
	if queryUser.CreatedAtFrom != nil && queryUser.CreatedAtTo != nil {
		where = append(where, " updated_at BETWEEN ? AND ?")
		args = append(args, queryUser.CreatedAtFrom, queryUser.CreatedAtTo)
	}
	if queryUser.UpdatedAtFrom != nil && queryUser.UpdatedAtTo != nil {
		where = append(where, " updated_at BETWEEN ? AND ?")
		args = append(args, queryUser.UpdatedAtFrom, queryUser.UpdatedAtTo)
	}

	if len(where) > 0 {
		queryString += "WHERE" + strings.Join(where, " AND") + "\n"
	}

	queryCountList := connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	}

	queryString += "ORDER BY _id ASC\n"

	queryUser.ParsePaging()
	queryString += "LIMIT ?, ?\n"

	args = append(args, queryUser.Offset)
	args = append(args, queryUser.Limit)

	queryList := connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	}

	resolveChan := make(chan interface{}, 2)
	rejectChan := make(chan error)

	go func() {
		data, err := factories.FindUser(&queryList)

		if err != nil {
			rejectChan <- err
		} else {
			resolveChan <- data
		}
	}()

	go func() {
		total, err := factories.CountUser(&queryCountList)

		if err != nil {
			rejectChan <- err
		}

		resolveChan <- total
	}()

	var data []*model.User
	var total int

	for i := 0; i < 2; i++ {
		select {
		case res := <-resolveChan:
			switch val := res.(type) {
			case []*model.User:
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
		Page:  *queryUser.Page,
		Limit: *queryUser.Limit,
	}, nil
}
