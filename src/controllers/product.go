package controllers

import (
	"strconv"

	"github.com/barrydev/api-3h-shop/src/actions"
	"github.com/barrydev/api-3h-shop/src/model"
	"github.com/gin-gonic/gin"
)

func GetListProduct(c *gin.Context) (interface{}, error) {
	var query model.QueryProduct

	err := c.ShouldBindQuery(&query)

	if err != nil {
		return nil, err
	}

	return actions.GetListProduct(&query)
}

func GetProductById(c *gin.Context) (interface{}, error) {
	productId, err := strconv.ParseInt(c.Param("productId"), 10, 64)

	if err != nil {
		return nil, err
	}

	return actions.GetProductById(productId)
}

func AddOldProductTag(c *gin.Context) (interface{}, error) {
	return actions.AddOldProductTag()
}

func GetProductItemByProductId(c *gin.Context) (interface{}, error) {
	productId, err := strconv.ParseInt(c.Param("productId"), 10, 64)

	if err != nil {
		return nil, err
	}

	return actions.GetProductItemByProductId(productId)
}

func InsertProductItemByProductId(c *gin.Context) (interface{}, error) {
	productId, err := strconv.ParseInt(c.Param("productId"), 10, 64)

	if err != nil {
		return nil, err
	}

	var insertProductItem model.BodyProductItem

	err = c.ShouldBindJSON(&insertProductItem)

	if err != nil {
		return nil, err
	}

	return actions.InsertProductItemByProductId(productId, &insertProductItem)
}

func InsertProduct(c *gin.Context) (interface{}, error) {
	var insertProduct model.BodyProduct

	err := c.ShouldBindJSON(&insertProduct)

	if err != nil {
		return nil, err
	}

	return actions.InsertProduct(&insertProduct)
}

func BulkInsertProduct(c *gin.Context) (interface{}, error) {
	var sliceBody model.SliceBodyProduct

	err := c.ShouldBindJSON(&sliceBody)

	if err != nil {
		return nil, err
	}

	return actions.BulkInsertProduct(&sliceBody)

	//return nil, nil
}

func BulkUpdateProduct(c *gin.Context) (interface{}, error) {
	var sliceBody model.SliceBodyProduct

	err := c.ShouldBindJSON(&sliceBody)

	if err != nil {
		return nil, err
	}

	return actions.BulkUpdateProduct(&sliceBody)

	//return nil, nil
}

func UpdateProduct(c *gin.Context) (interface{}, error) {
	productId, err := strconv.ParseInt(c.Param("productId"), 10, 64)

	if err != nil {
		return nil, err
	}

	var body model.BodyProduct

	err = c.ShouldBindJSON(&body)

	if err != nil {
		return nil, err
	}

	return actions.UpdateProduct(productId, &body)
}
