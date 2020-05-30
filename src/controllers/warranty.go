package controllers

import (
	"strconv"

	"github.com/barrydev/api-3h-shop/src/actions"
	"github.com/barrydev/api-3h-shop/src/model"
	"github.com/gin-gonic/gin"
)

func GetListWarranty(c *gin.Context) (interface{}, error) {
	var query model.QueryWarranty

	err := c.ShouldBindQuery(&query)

	if err != nil {
		return nil, err
	}

	return actions.GetListWarranty(&query)
}

func GetWarrantyById(c *gin.Context) (interface{}, error) {
	warrantyId, err := strconv.ParseInt(c.Param("warrantyId"), 10, 64)

	if err != nil {
		return nil, err
	}

	return actions.GetWarrantyById(warrantyId)
}

func InsertWarranty(c *gin.Context) (interface{}, error) {
	var insertWarranty model.BodyWarranty

	err := c.ShouldBindJSON(&insertWarranty)

	if err != nil {
		return nil, err
	}

	return actions.InsertWarranty(&insertWarranty)
}

func UpdateWarranty(c *gin.Context) (interface{}, error) {
	warrantyId, err := strconv.ParseInt(c.Param("warrantyId"), 10, 64)

	if err != nil {
		return nil, err
	}

	var body model.BodyWarranty

	err = c.ShouldBindJSON(&body)

	if err != nil {
		return nil, err
	}

	return actions.UpdateWarranty(warrantyId, &body)
}
