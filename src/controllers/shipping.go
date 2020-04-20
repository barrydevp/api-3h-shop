package controllers

import (
	"github.com/barrydev/api-3h-shop/src/actions"
	"github.com/barrydev/api-3h-shop/src/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetListShipping(c *gin.Context) (interface{}, error) {
	var query model.QueryShipping

	err := c.ShouldBindQuery(&query)

	if err != nil {
		return nil, err
	}

	return actions.GetListShipping(&query)
}

func GetShippingById(c *gin.Context) (interface{}, error) {
	shippingId, err := strconv.ParseInt(c.Param("shippingId"), 10, 64)

	if err != nil {
		return nil, err
	}

	return actions.GetShippingById(shippingId)
}

func InsertShipping(c *gin.Context) (interface{}, error) {
	var insertShipping model.BodyShipping

	err := c.ShouldBindJSON(&insertShipping)

	if err != nil {
		return nil, err
	}

	return actions.InsertShipping(&insertShipping)
}

func UpdateShipping(c *gin.Context) (interface{}, error) {
	shippingId, err := strconv.ParseInt(c.Param("shippingId"), 10, 64)

	if err != nil {
		return nil, err
	}

	var body model.BodyShipping

	err = c.ShouldBindJSON(&body)

	if err != nil {
		return nil, err
	}

	return actions.UpdateShipping(shippingId, &body)
}