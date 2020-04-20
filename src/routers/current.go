package routers

import (
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/controllers"
	"github.com/gin-gonic/gin"
)

func BindOrder(router *gin.RouterGroup) {

	router.GET("/", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetListOrder).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.GET("/:orderId", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetOrderById).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.GET("/:orderId/items", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetOrderItemByOrderId).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.GET("/:orderId/shipping", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetShippingByOrderId).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.POST("/:orderId/items", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.InsertOrderItemByOrderId).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.POST("/:orderId/shipping", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.InsertShippingByOrderId).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.POST("/", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.InsertOrder).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.POST("/:orderId/update", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.UpdateOrder).Then(response.SendSuccess).Catch(response.SendError)
	})
}
