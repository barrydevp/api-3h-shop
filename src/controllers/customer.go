package controllers

import (
	"github.com/barrydev/api-3h-shop/src/actions"
	"github.com/barrydev/api-3h-shop/src/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetListCustomer(c *gin.Context) (interface{}, error) {
	var query model.QueryCustomer

	err := c.ShouldBindQuery(&query)

	if err != nil {
		return nil, err
	}

	return actions.GetListCustomer(&query)
}

func GetCustomerById(c *gin.Context) (interface{}, error) {
	customerId, err := strconv.ParseInt(c.Param("customerId"), 10, 64)

	if err != nil {
		return nil, err
	}

	return actions.GetCustomerById(customerId)
}

func InsertCustomer(c *gin.Context) (interface{}, error) {
	var insertCustomer model.BodyCustomer

	err := c.ShouldBindJSON(&insertCustomer)

	if err != nil {
		return nil, err
	}

	return actions.InsertCustomer(&insertCustomer)
}

func UpdateCustomer(c *gin.Context) (interface{}, error) {
	customerId, err := strconv.ParseInt(c.Param("customerId"), 10, 64)

	if err != nil {
		return nil, err
	}

	var body model.BodyCustomer

	err = c.ShouldBindJSON(&body)

	if err != nil {
		return nil, err
	}

	return actions.UpdateCustomer(customerId, &body)
}