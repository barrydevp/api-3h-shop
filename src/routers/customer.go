package routers

import (
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/controllers"
	"github.com/gin-gonic/gin"
)

func BindCustomer(router *gin.RouterGroup) {

	router.GET("", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetListCustomer).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.GET("/:customerId", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetCustomerById).Then(response.SendSuccess).Catch(response.SendError)
	})

	// router.POST("", func(c *gin.Context) {
	// 	handle := response.Handle{Context: c}

	// 	handle.Try(controllers.InsertCustomer).Then(response.SendSuccess).Catch(response.SendError)
	// })

	// router.POST("/:customerId/update", func(c *gin.Context) {
	// 	handle := response.Handle{Context: c}

	// 	handle.Try(controllers.UpdateCustomer).Then(response.SendSuccess).Catch(response.SendError)
	// })
}

func BindCustomerAdmin(router *gin.RouterGroup) {

	router.POST("", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.InsertCustomer).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.POST("/:customerId/update", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.UpdateCustomer).Then(response.SendSuccess).Catch(response.SendError)
	})
}
