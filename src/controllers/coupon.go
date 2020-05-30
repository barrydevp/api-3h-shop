package controllers

import (
	"strconv"

	"github.com/barrydev/api-3h-shop/src/actions"
	"github.com/barrydev/api-3h-shop/src/model"
	"github.com/gin-gonic/gin"
)

func GetListCoupon(c *gin.Context) (interface{}, error) {
	var query model.QueryCoupon

	err := c.ShouldBindQuery(&query)

	if err != nil {
		return nil, err
	}

	return actions.GetListCoupon(&query)
}

func GetCouponById(c *gin.Context) (interface{}, error) {
	couponId, err := strconv.ParseInt(c.Param("couponId"), 10, 64)

	if err != nil {
		return nil, err
	}

	return actions.GetCouponById(couponId)
}

func InsertCoupon(c *gin.Context) (interface{}, error) {
	var insertCoupon model.BodyCoupon

	err := c.ShouldBindJSON(&insertCoupon)

	if err != nil {
		return nil, err
	}

	return actions.InsertCoupon(&insertCoupon)
}

func UpdateCoupon(c *gin.Context) (interface{}, error) {
	couponId, err := strconv.ParseInt(c.Param("couponId"), 10, 64)

	if err != nil {
		return nil, err
	}

	var body model.BodyCoupon

	err = c.ShouldBindJSON(&body)

	if err != nil {
		return nil, err
	}

	return actions.UpdateCoupon(couponId, &body)
}
