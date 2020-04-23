package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
	"log"
	"strings"
)

func BulkUpdateProduct(sliceBody *model.SliceBodyProduct) (interface{}, error) {

	for index, body := range sliceBody.Data{
		log.Println("Start on index - ", index)
		queryString := ""

		var args []interface{}

		var set []string

		if body.Name == nil {
			return nil, errors.New("product's name is required")
		} else {
			set = append(set, " name=?")
			args = append(args, body.Name)
		}

		if body.ImagePath != nil {
			set = append(set, " image_path=?")
			args = append(args, body.ImagePath)
		}

		if body.Description != nil {
			set = append(set, " description=?")
			args = append(args, body.Description)
		}

		if len(set) > 0 {
			queryString += "SET" + strings.Join(set, ",") + "\n"
		} else {
			return nil, errors.New("invalid body")
		}

		queryString += "WHERE name=?"
		args = append(args, body.Name)

		log.Println(queryString)

		rowEffected, err := factories.UpdateProduct(&connect.QueryMySQL{
			QueryString: queryString,
			Args:        args,
		})

		if err != nil {
			return nil, err
		}

		if rowEffected == nil {
			return nil, errors.New("insert error")
		}

		log.Println("Completed on index - ", index, "rowEffected - ", rowEffected)
	}


	return nil, nil
}
