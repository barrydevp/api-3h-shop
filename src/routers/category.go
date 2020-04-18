package routers

import (
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/controllers"
	"github.com/gin-gonic/gin"
)

func BindCategory(router *gin.RouterGroup) {

	router.GET("/", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetListCategory).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.GET("/:categoryId", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetOneCategory).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.POST("/", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.InsertCategory).Then(response.SendSuccess).Catch(response.SendError)
	})
}