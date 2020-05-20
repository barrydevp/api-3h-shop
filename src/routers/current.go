package routers

import (
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/controllers"
	"github.com/gin-gonic/gin"
)

func BindCurrent(router *gin.RouterGroup) {

	router.GET("/order", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetCurrentOrderV2).Then(response.SendSuccess).Catch(response.SendError)
	})
}

func BindCurrentAuth(router *gin.RouterGroup) {

	router.GET("/order", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetCurrentOrderAuth).Then(response.SendSuccess).Catch(response.SendError)
	})
}
