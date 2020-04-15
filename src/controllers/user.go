package controllers

import (
	"github.com/barrydev/api-3h-shop/src/actions"
	"github.com/gin-gonic/gin"
	"log"
)

func GetListUser(c *gin.Context) (interface{}, error) {
	log.Println(c.Request.URL.Query())
	return actions.GetListUser()
}
