package controllers

import (
	"github.com/barrydev/api-3h-shop/src/actions"
	"github.com/barrydev/api-3h-shop/src/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetListOrderItem(c *gin.Context) (interface{}, error) {
	var query model.QueryOrderItem

	err := c.ShouldBindQuery(&query)

	if err != nil {
		return nil, err
	}

	return actions.GetListOrderItem(&query)
}

func GetOrderItemById(c *gin.Context) (interface{}, error) {
	orderItemId, err := strconv.ParseInt(c.Param("orderItemId"), 10, 64)

	if err != nil {
		return nil, err
	}

	return actions.GetOrderItemById(orderItemId)
}

func InsertOrderItem(c *gin.Context) (interface{}, error) {
	var insertOrderItem model.BodyOrderItem

	err := c.ShouldBindJSON(&insertOrderItem)

	if err != nil {
		return nil, err
	}

	return actions.InsertOrderItem(&insertOrderItem)
}

func UpdateOrderItem(c *gin.Context) (interface{}, error) {
	orderItemId, err := strconv.ParseInt(c.Param("orderItemId"), 10, 64)

	if err != nil {
		return nil, err
	}

	var body model.BodyOrderItem

	err = c.ShouldBindJSON(&body)

	if err != nil {
		return nil, err
	}

	return actions.UpdateOrderItem(orderItemId, &body)
}

func DeleteOrderItemById(c *gin.Context) (interface{}, error) {
	orderItemId, err := strconv.ParseInt(c.Param("orderItemId"), 10, 64)

	if err != nil {
		return nil, err
	}

	return actions.DeleteOrderItem(orderItemId)
}

func DecreaseQuantityOrderItem(c *gin.Context) (interface{}, error) {
	orderItemId, err := strconv.ParseInt(c.Param("orderItemId"), 10, 64)

	if err != nil {
		return nil, err
	}

	var body model.BodyOrderItem

	err = c.ShouldBindJSON(&body)

	if err != nil {
		return nil, err
	}

	return actions.DecreaseQuantityOrderItem(orderItemId, &body)
}

func IncreaseQuantityOrderItem(c *gin.Context) (interface{}, error) {
	orderItemId, err := strconv.ParseInt(c.Param("orderItemId"), 10, 64)

	if err != nil {
		return nil, err
	}

	var body model.BodyOrderItem

	err = c.ShouldBindJSON(&body)

	if err != nil {
		return nil, err
	}

	return actions.IncreaseQuantityOrderItem(orderItemId, &body)
}
