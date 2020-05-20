package controllers

import (
	"errors"
	"strconv"

	"github.com/barrydev/api-3h-shop/src/actions"
	"github.com/barrydev/api-3h-shop/src/constants"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
	"github.com/gin-gonic/gin"
)

func GetCurrentOrderAuth(c *gin.Context) (interface{}, error) {
	var current model.Current

	session, err := c.Cookie("3h.session")

	if err != nil {
		current.Session = nil
	} else {
		current.Session = &session
	}

	payload, ok := c.Get("payload_token")

	if !ok {
		return nil, errors.New("invalid payload_token")
	}

	payloadToken, ok := payload.(*factories.AccessTokenClaims)

	if !ok {
		return nil, errors.New("invalid payload_token")
	}

	res, err := actions.GetCurrentOrderAuth(&current, payloadToken)

	if err != nil {
		return nil, err
	}

	c.SetCookie("3h.session", *current.Session, 1000*30*24*3600, "/", constants.PRIMARY_HOST, false, true)
	//c.SetCookie("3h.session", *current.Session, 1000*30*24*3600, "/", "web-3h-shop.herokuapp.com", false, true)
	return res, nil
}

func GetCurrentOrderV2(c *gin.Context) (interface{}, error) {
	var current model.Current

	session, err := c.Cookie("3h.session")

	if err != nil {
		current.Session = nil
	} else {
		current.Session = &session
	}

	res, err := actions.GetCurrentOrderV2(&current)

	if err != nil {
		return nil, err
	}

	c.SetCookie("3h.session", *current.Session, 1000*30*24*3600, "/", constants.PRIMARY_HOST, false, true)
	//c.SetCookie("3h.session", *current.Session, 1000*30*24*3600, "/", "web-3h-shop.herokuapp.com", false, true)
	return res, nil
}

func GetCurrentOrder(c *gin.Context) (interface{}, error) {
	var current model.Current

	session, err := c.Cookie("3h.session")

	if err != nil {
		current.Session = nil
	} else {
		current.Session = &session
	}

	res, err := actions.GetCurrentOrder(&current)

	if err != nil {
		return nil, err
	}

	c.SetCookie("3h.session", *current.Session, 1000*30*24*3600, "/", constants.PRIMARY_HOST, false, true)
	//c.SetCookie("3h.session", *current.Session, 1000*30*24*3600, "/", "web-3h-shop.herokuapp.com", false, true)
	return res, nil
}

func GetOrderItemByaOrderId(c *gin.Context) (interface{}, error) {
	orderId, err := strconv.ParseInt(c.Param("orderId"), 10, 64)

	if err != nil {
		return nil, err
	}

	return actions.GetOrderItemByOrderId(orderId)
}

func InsertOrderItemByOrderIda(c *gin.Context) (interface{}, error) {
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

func GetShippingByOrderId1(c *gin.Context) (interface{}, error) {
	orderId, err := strconv.ParseInt(c.Param("orderId"), 10, 64)

	if err != nil {
		return nil, err
	}

	return actions.GetShippingByOrderId(orderId)
}
