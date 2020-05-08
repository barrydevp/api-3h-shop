package routers

import (
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/controllers"
	"github.com/gin-gonic/gin"
)

func BindShipping(router *gin.RouterGroup) {

	router.GET("", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetListShipping).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.GET("/:shippingId", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetShippingById).Then(response.SendSuccess).Catch(response.SendError)
	})

	// router.POST("", func(c *gin.Context) {
	// 	handle := response.Handle{Context: c}

	// 	handle.Try(controllers.InsertShipping).Then(response.SendSuccess).Catch(response.SendError)
	// })

	// router.POST("/:shippingId/update", func(c *gin.Context) {
	// 	handle := response.Handle{Context: c}

	// 	handle.Try(controllers.UpdateShipping).Then(response.SendSuccess).Catch(response.SendError)
	// })
}
