package controllers

import (
	"github.com/barrydev/api-3h-shop/src/actions"
	"github.com/barrydev/api-3h-shop/src/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetListProductItem(c *gin.Context) (interface{}, error) {
	var query model.QueryProductItem

	err := c.ShouldBindQuery(&query)

	if err != nil {
		return nil, err
	}

	return actions.GetListProductItem(&query)
}

func GetProductItemById(c *gin.Context) (interface{}, error) {
	productItemId, err := strconv.ParseInt(c.Param("productItemId"), 10, 64)

	if err != nil {
		return nil, err
	}

	return actions.GetProductItemById(productItemId)
}

func InsertProductItem(c *gin.Context) (interface{}, error) {
	var insertProductItem model.BodyProductItem

	err := c.ShouldBindJSON(&insertProductItem)

	if err != nil {
		return nil, err
	}

	return actions.InsertProductItem(&insertProductItem)
}

func UpdateProductItem(c *gin.Context) (interface{}, error) {
	productItemId, err := strconv.ParseInt(c.Param("productItemId"), 10, 64)

	if err != nil {
		return nil, err
	}

	var body model.BodyProductItem

	err = c.ShouldBindJSON(&body)

	if err != nil {
		return nil, err
	}

	return actions.UpdateProductItem(productItemId, &body)
}