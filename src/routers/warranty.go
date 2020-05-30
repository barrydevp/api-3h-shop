package routers

import (
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/controllers"
	"github.com/gin-gonic/gin"
)

func BindWarranty(router *gin.RouterGroup) {

	router.GET("", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetListWarranty).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.GET("/:warrantyId", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetWarrantyById).Then(response.SendSuccess).Catch(response.SendError)
	})

}

func BindWarrantyAdmin(router *gin.RouterGroup) {

	router.POST("", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.InsertWarranty).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.POST("/:warrantyId/update", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.UpdateWarranty).Then(response.SendSuccess).Catch(response.SendError)
	})
}
