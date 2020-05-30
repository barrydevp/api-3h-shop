package controllers

import (
	"strconv"

	"github.com/barrydev/api-3h-shop/src/actions"
	"github.com/barrydev/api-3h-shop/src/model"
	"github.com/gin-gonic/gin"
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

func GetOrderItemByOrderId(c *gin.Context) (interface{}, error) {
	orderId, err := strconv.ParseInt(c.Param("orderId"), 10, 64)

	if err != nil {
		return nil, err
	}

	return actions.GetOrderItemByOrderId(orderId)
}

func InsertOrderItemByOrderId(c *gin.Context) (interface{}, error) {
	orderId, err := strconv.ParseInt(c.Param("orderId"), 10, 64)

	if err != nil {
		return nil, err
	}

	var insertOrderItem model.BodyOrderItem

	err = c.ShouldBindJSON(&insertOrderItem)

	if err != nil {
		return nil, err
	}

	return actions.InsertOrderItemByOrderId(orderId, &insertOrderItem)
}

func ApplyCoupon(c *gin.Context) (interface{}, error) {
	orderId, err := strconv.ParseInt(c.Param("orderId"), 10, 64)

	if err != nil {
		return nil, err
	}

	var bodyCoupon model.BodyCoupon

	err = c.ShouldBindJSON(&bodyCoupon)

	if err != nil {
		return nil, err
	}

	return actions.ApplyCoupon(orderId, &bodyCoupon)
}

func RemoveOrderCoupon(c *gin.Context) (interface{}, error) {
	orderId, err := strconv.ParseInt(c.Param("orderId"), 10, 64)

	if err != nil {
		return nil, err
	}

	return actions.RemoveOrderCoupon(orderId)
}

func GetOrderCoupon(c *gin.Context) (interface{}, error) {
	orderId, err := strconv.ParseInt(c.Param("orderId"), 10, 64)

	if err != nil {
		return nil, err
	}

	return actions.GetOrderCoupon(orderId)
}

func GetShippingByOrderId(c *gin.Context) (interface{}, error) {
	orderId, err := strconv.ParseInt(c.Param("orderId"), 10, 64)

	if err != nil {
		return nil, err
	}

	return actions.GetShippingByOrderId(orderId)
}

func InsertShippingByOrderId(c *gin.Context) (interface{}, error) {
	orderId, err := strconv.ParseInt(c.Param("orderId"), 10, 64)

	if err != nil {
		return nil, err
	}

	var insertShipping model.BodyShipping

	err = c.ShouldBindJSON(&insertShipping)

	if err != nil {
		return nil, err
	}

	return actions.InsertShippingByOrderId(orderId, &insertShipping)
}

func CheckoutOrder(c *gin.Context) (interface{}, error) {
	orderId, err := strconv.ParseInt(c.Param("orderId"), 10, 64)

	if err != nil {
		return nil, err
	}

	var checkoutBody model.BodyCheckoutOrder

	err = c.ShouldBindJSON(&checkoutBody)

	if err != nil {
		return nil, err
	}

	return actions.CheckoutOrder(orderId, &checkoutBody)
}

func InsertOrder(c *gin.Context) (interface{}, error) {
	var insertOrder model.BodyOrder

	err := c.ShouldBindJSON(&insertOrder)

	if err != nil {
		return nil, err
	}

	return actions.InsertOrder(&insertOrder)
}

func GetOrderCustomerByOrderId(c *gin.Context) (interface{}, error) {
	orderId, err := strconv.ParseInt(c.Param("orderId"), 10, 64)

	if err != nil {
		return nil, err
	}

	return actions.GetOrderCustomerByOrderId(orderId)
}

func UpdateOrderCustomer(c *gin.Context) (interface{}, error) {
	orderId, err := strconv.ParseInt(c.Param("orderId"), 10, 64)

	if err != nil {
		return nil, err
	}
	var body model.BodyCustomer

	err = c.ShouldBindJSON(&body)

	if err != nil {
		return nil, err
	}

	return actions.UpdateOrderCustomer(orderId, &body)
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

func ChangeOrderNoteByOrderId(c *gin.Context) (interface{}, error) {
	orderId, err := strconv.ParseInt(c.Param("orderId"), 10, 64)

	if err != nil {
		return nil, err
	}

	var body model.BodyOrder

	err = c.ShouldBindJSON(&body)

	if err != nil {
		return nil, err
	}

	return actions.ChangeOrderNoteByOrderId(orderId, &body)
}

func ChangeOrderFulfilmentStatusByOrderId(c *gin.Context) (interface{}, error) {
	orderId, err := strconv.ParseInt(c.Param("orderId"), 10, 64)

	if err != nil {
		return nil, err
	}

	var body model.BodyOrder

	err = c.ShouldBindJSON(&body)

	if err != nil {
		return nil, err
	}

	return actions.ChangeOrderFulfilmentStatusByOrderId(orderId, &body)
}

func ChangeOrderPaymentStatusByOrderId(c *gin.Context) (interface{}, error) {
	orderId, err := strconv.ParseInt(c.Param("orderId"), 10, 64)

	if err != nil {
		return nil, err
	}

	var body model.BodyOrder

	err = c.ShouldBindJSON(&body)

	if err != nil {
		return nil, err
	}

	return actions.ChangeOrderPaymentStatusByOrderId(orderId, &body)
}

func MarkOrderPaid(c *gin.Context) (interface{}, error) {
	orderId, err := strconv.ParseInt(c.Param("orderId"), 10, 64)

	if err != nil {
		return nil, err
	}

	return actions.MarkOrderPaid(orderId)
}
