package controllers

import (
	"github.com/barrydev/api-3h-shop/src/actions"
	"github.com/barrydev/api-3h-shop/src/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetListOrder(c *gin.Context) (interface{}, error) {
	var query model.QueryOrder

	err := c.ShouldBindQuery(&query)

	if err != nil {
		return nil, err
	}

	return actions.GetListOrder(&query)
}

func GetOrderById(c *gin.Context) (interface{}, error) {
	orderId, err := strconv.ParseInt(c.Param("orderId"), 10, 64)

	if err != nil {
		return nil, err
	}

	return actions.GetOrderById(orderId)
}

func InsertOrder(c *gin.Context) (interface{}, error) {
	var insertOrder model.BodyOrder

	err := c.ShouldBindJSON(&insertOrder)

	if err != nil {
		return nil, err
	}

	return actions.InsertOrder(&insertOrder)
}

func UpdateOrder(c *gin.Context) (interface{}, error) {
	orderId, err := strconv.ParseInt(c.Param("orderId"), 10, 64)

	if err != nil {
		return nil, err
	}

	var body model.BodyOrder

	err = c.ShouldBindJSON(&body)

	if err != nil {
		return nil, err
	}

	return actions.UpdateOrder(orderId, &body)
}