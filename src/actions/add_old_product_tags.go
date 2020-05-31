package actions

import (
	"log"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func _addTags(list []*model.Product) {
	for index, _ := range list {

		// tagsFromName := strings.Split(*list[index].Name, " ")

		// log.Println(tagsFromName)

		// originalLen := len(tagsFromName)

		// for i := 0; i < originalLen; i++ {
		// 	current := tagsFromName[i]
		// 	for j := i + 1; j < originalLen; j++ {
		// 		current += tagsFromName[j]
		// 		tagsFromName = append(tagsFromName, current)
		// 	}
		// }

		_, err := factories.UpdateProduct(&connect.QueryMySQL{
			QueryString: "SET tags=COALESCE((SELECT CONCAT(c.name, ',', (SELECT pc.name FROM categories pc WHERE _id=c.parent_id), ',') FROM categories c WHERE _id=?), '') WHERE _id=?",
			Args:        []interface{}{list[index].CategoryId, list[index].Id},
		})

		if err != nil {
			log.Println(err)
		} else {
			log.Println("Update success tags of _id= ", *list[index].Id)
		}
	}
}

func _workerAction(page, limit int) bool {
	list, err := factories.FindProduct(&connect.QueryMySQL{
		QueryString: "ORDER BY _id ASC \n LIMIT ?, ?",
		Args:        []interface{}{page * limit, limit},
	})

	if err != nil {
		return false
	}

	if list == nil {
		return false
	}

	_addTags(list)

	return true
}

func AddOldProductTag() (bool, error) {
	go func() {
		limit := 20
		page := 0

		for ; _workerAction(page, limit); page++ {

		}
	}()

	return true, nil
}
