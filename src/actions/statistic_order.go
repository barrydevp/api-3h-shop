package actions

import (
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func StatisticOrder(query *model.QueryStatisticOrder) (*model.StatisticOrder, error) {
	queryString := ""
	var args []interface{}

	var where []string

	if query.CreatedAtFrom != nil && query.CreatedAtTo != nil {
		where = append(where, " created_at BETWEEN ? AND ?")
		args = append(args, query.CreatedAtFrom, query.CreatedAtTo)
	} else {
		if query.CreatedAtFrom != nil {
			where = append(where, " created_at>=?")
			args = append(args, query.CreatedAtFrom)
		}

		if query.CreatedAtTo != nil {
			where = append(where, " created_at<=?")
			args = append(args, query.CreatedAtTo)
		}

	}

	if len(where) > 0 {
		queryString += "WHERE" + strings.Join(where, " AND") + "\n"
	}

	statistic, err := factories.StatisticOrder(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	return statistic, nil
}
