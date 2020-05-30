package routers

import (
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/controllers"
	"github.com/gin-gonic/gin"
)

func BindCoupon(router *gin.RouterGroup) {

	router.GET("", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetListCoupon).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.GET("/:couponId", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetCouponById).Then(response.SendSuccess).Catch(response.SendError)
	})

}

func BindCouponAdmin(router *gin.RouterGroup) {

	router.POST("", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.InsertCoupon).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.POST("/:couponId/update", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.UpdateCoupon).Then(response.SendSuccess).Catch(response.SendError)
	})
}
