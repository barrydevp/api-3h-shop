package actions

import (
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetListCategory(queryCategory *model.QueryCategory) (*response.DataList, error) {
	queryString := ""
	var args []interface{}

	var where []string

	if queryCategory.Id != nil {
		where = append(where, " _id=?")
		args = append(args, queryCategory.Id)
	}
	if queryCategory.Name != nil {
		where = append(where, " name LIKE ?")
		args = append(args, "%"+*queryCategory.Name+"%")
	}
	if queryCategory.ParentId != nil {
		if *queryCategory.ParentId != 0 {
			if *queryCategory.ParentId == -1 {
				where = append(where, " parent_id IS NULL")
			} else {
				where = append(where, " parent_id=?")
				args = append(args, queryCategory.ParentId)
			}
		}
	}
	if queryCategory.Status != nil {
		where = append(where, " status=?")
		args = append(args, queryCategory.Status)
	}
	if queryCategory.UpdatedAtFrom != nil && queryCategory.UpdatedAtTo != nil {
		where = append(where, " updated_at BETWEEN ? AND ?")
		args = append(args, queryCategory.UpdatedAtFrom, queryCategory.UpdatedAtTo)
	}

	if len(where) > 0 {
		queryString += "WHERE" + strings.Join(where, " AND") + "\n"
	}

	queryCountList := connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	}

	queryString += "ORDER BY _id ASC\n"

	queryCategory.ParsePaging()
	queryString += "LIMIT ?, ?\n"

	args = append(args, queryCategory.Offset)
	args = append(args, queryCategory.Limit)

	queryList := connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	}

	resolveChan := make(chan interface{}, 2)
	rejectChan := make(chan error)

	go func() {
		data, err := factories.FindCategory(&queryList)

		if err != nil {
			rejectChan <- err
		} else {
			resolveChan <- data
		}
	}()

	go func() {
		total, err := factories.CountCategory(&queryCountList)

		if err != nil {
			rejectChan <- err
		}

		resolveChan <- total
	}()

	var data []*model.Category
	var total int

	for i := 0; i < 2; i++ {
		select {
		case res := <-resolveChan:
			switch val := res.(type) {
			case []*model.Category:
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
		Page:  *queryCategory.Page,
		Limit: *queryCategory.Limit,
	}, nil
}
