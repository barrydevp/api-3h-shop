package factories

import (
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindOrderItemJoinProduct(query *connect.QueryMySQL) ([]*model.OrderItemJoinProduct, error) {
	connection := connections.Mysql.GetConnection()

	queryString := `
		SELECT
			orderitems._id, orderitems.product_id, orderitems.order_id, orderitems.quantity, orderitems.status, orderitems.created_at, orderitems.updated_at,
			products._id, products.category_id, products.name, products.out_price, products.discount, products.image_path, products.created_at, products.updated_at 
		FROM (SELECT
			order_items._id, order_items.product_id, order_items.product_item_id, order_items.order_id, order_items.quantity, order_items.status, order_items.created_at, order_items.updated_at
		FROM order_items
	`
	var args []interface{}

	if query != nil {
		queryString += query.QueryString + ") AS orderitems\n"
		args = query.Args
	}

	queryString += `
		INNER JOIN products ON orderitems.product_id = products._id	
	`

	stmt, err := connection.Prepare(queryString)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var listOrderItemJoinProducts []*model.OrderItemJoinProduct

	for rows.Next() {
		_orderItem := model.OrderItem{}
		_product := model.Product{}
		_orderItemJoinProduct := model.OrderItemJoinProduct{
			OrderItem: &_orderItem,
			Product: &_product,
		}

		err = rows.Scan(
			&_orderItem.RawId,
			&_orderItem.RawProductId,
			&_orderItem.RawOrderId,
			&_orderItem.RawQuantity,
			&_orderItem.RawStatus,
			&_orderItem.RawCreatedAt,
			&_orderItem.RawUpdatedAt,
			&_product.RawId,
			&_product.RawCategoryId,
			&_product.RawName,
			&_product.RawOutPrice,
			&_product.RawDiscount,
			&_product.RawImagePath,
			&_product.RawCreatedAt,
			&_product.RawUpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		_orderItemJoinProduct.FillResponse()

		listOrderItemJoinProducts = append(listOrderItemJoinProducts, &_orderItemJoinProduct)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return listOrderItemJoinProducts, nil
}
