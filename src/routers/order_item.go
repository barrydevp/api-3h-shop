package routers

import (
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/controllers"
	"github.com/gin-gonic/gin"
)

func BindOrderItem(router *gin.RouterGroup) {

	router.GET("/", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetListOrderItem).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.GET("/:orderItemId", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetOrderItemById).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.POST("/", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.InsertOrderItem).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.POST("/:orderItemId/update", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.UpdateOrderItem).Then(response.SendSuccess).Catch(response.SendError)
	})
}
