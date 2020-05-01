package actions

import (
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
	"strings"
)

func GetListProduct(queryProduct *model.QueryProduct) (*response.DataList, error) {
	queryString := ""
	var args []interface{}

	var where []string

	if queryProduct.Id != nil {
		where = append(where, " _id=?")
		args = append(args, queryProduct.Id)
	}
	if queryProduct.CategoryId != nil {
		if *queryProduct.CategoryId != 0 {
			if *queryProduct.CategoryId == -1 {
				where = append(where, " category_id IS NULL")
			} else {
				where = append(where, " category_id=?")
				args = append(args, queryProduct.CategoryId)
			}
		}
	}
	if queryProduct.CategoryParentId != nil {
		if *queryProduct.CategoryParentId != 0 {
			where = append(where, " EXISTS (SELECT _id FROM categories WHERE categories.parent_id=? AND categories._id=products.category_id)")
			args = append(args, queryProduct.CategoryParentId)
		}
	}
	if queryProduct.Name != nil {
		where = append(where, " name LIKE ?")
		args = append(args, "%"+*queryProduct.Name+"%")
	}
	if queryProduct.StartOutPrice != nil && *queryProduct.StartOutPrice > 0 {
		where = append(where, " out_price >= ?")
		args = append(args, queryProduct.StartOutPrice)
	}
	if queryProduct.EndOutPrice != nil && *queryProduct.EndOutPrice > 0 {
		where = append(where, " out_price <= ?")
		args = append(args, queryProduct.EndOutPrice)
	}
	if queryProduct.CreatedAtFrom != nil && queryProduct.CreatedAtTo != nil {
		where = append(where, " created_at BETWEEN ? AND ?")
		args = append(args, queryProduct.CreatedAtFrom, queryProduct.CreatedAtTo)
	}
	if queryProduct.UpdatedAtFrom != nil && queryProduct.UpdatedAtTo != nil {
		where = append(where, " updated_at BETWEEN ? AND ?")
		args = append(args, queryProduct.UpdatedAtFrom, queryProduct.UpdatedAtTo)
	}

	if len(where) > 0 {
		queryString += "WHERE" + strings.Join(where, " AND") + "\n"
	}

	queryCountList := connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	}

	queryProduct.ParseSort()
	queryString += "ORDER BY " + *queryProduct.OrderBy + "\n"

	queryProduct.ParsePaging()
	queryString += "LIMIT ?, ?\n"

	args = append(args, queryProduct.Offset)
	args = append(args, queryProduct.Limit)

	queryList := connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	}

	resolveChan := make(chan interface{}, 2)
	rejectChan := make(chan error)

	go func() {
		data, err := factories.FindProduct(&queryList)

		if err != nil {
			rejectChan <- err
		} else {
			resolveChan <- data
		}
	}()

	go func() {
		total, err := factories.CountProduct(&queryCountList)

		if err != nil {
			rejectChan <- err
		}

		resolveChan <- total
	}()

	var data []*model.Product
	var total int

	for i := 0; i < 2; i++ {
		select {
		case res := <-resolveChan:
			switch val := res.(type) {
			case []*model.Product:
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
		Page:  *queryProduct.Page,
		Limit: *queryProduct.Limit,
	}, nil
}
