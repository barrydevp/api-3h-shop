package controllers

import (
	"github.com/barrydev/api-3h-shop/src/actions"
	"github.com/barrydev/api-3h-shop/src/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetListCategory(c *gin.Context) (interface{}, error) {
	var query model.QueryCategory

	err := c.ShouldBindQuery(&query)

	if err != nil {
		return nil, err
	}

	return actions.GetListCategory(query)
}

func GetOneCategory(c *gin.Context) (interface{}, error) {
	categoryId, err := strconv.ParseInt(c.Param("categoryId"), 10, 64)

	if err != nil {
		return nil, err
	}

	return actions.GetOneCategory(&categoryId)
}

func InsertCategory(c *gin.Context) (interface{}, error) {
	var insertCategory model.BodyCategory

	err := c.ShouldBindJSON(&insertCategory)

	if err != nil {
		return nil, err
	}

	return actions.InsertCategory(&insertCategory)
}

func GetAllCategoryTree(c *gin.Context) (interface{}, error) {
	return actions.GetAllCategoryTree()
}
