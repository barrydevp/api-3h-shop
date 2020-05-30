package controllers

import (
	"github.com/barrydev/api-3h-shop/src/actions"
	"github.com/barrydev/api-3h-shop/src/model"
	"github.com/gin-gonic/gin"
)

func StatisticOrder(c *gin.Context) (interface{}, error) {
	var query model.QueryStatisticOrder

	err := c.ShouldBindQuery(&query)

	if err != nil {
		return nil, err
	}

	return actions.StatisticOrder(&query)
}
