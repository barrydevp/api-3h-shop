package routers

import (
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/controllers"
	"github.com/gin-gonic/gin"
)

func BindProduct(router *gin.RouterGroup) {

	router.GET("", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetListProduct).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.GET("/:productId", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetProductById).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.GET("/:productId/items", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetProductItemByProductId).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.POST("/:productId/items", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.InsertProductItemByProductId).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.POST("", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.InsertProduct).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.POST("/:productId/update", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.UpdateProduct).Then(response.SendSuccess).Catch(response.SendError)
	})
}
