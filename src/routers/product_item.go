package routers

import (
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/controllers"
	"github.com/gin-gonic/gin"
)

func BindProductItem(router *gin.RouterGroup) {

	router.GET("/", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetListProductItem).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.GET("/:productItemId", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetProductItemById).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.POST("/", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.InsertProductItem).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.POST("/:productItemId/update", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.UpdateProductItem).Then(response.SendSuccess).Catch(response.SendError)
	})
}
