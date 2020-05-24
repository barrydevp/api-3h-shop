package routers

import (
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/controllers"
	"github.com/gin-gonic/gin"
)

func BindUser(router *gin.RouterGroup) {

	router.POST("/register", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.RegisterUser).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.POST("/authenticate", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.AuthenticateUser).Then(response.SendSuccess).Catch(response.SendError)
	})
}

func BindUserAuth(router *gin.RouterGroup) {

	router.GET("", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetListUser).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.GET("/:userId", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetUserById).Then(response.SendSuccess).Catch(response.SendError)
	})
}

func BindUserAdmin(router *gin.RouterGroup) {

	router.GET("", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetListUser).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.GET("/:userId", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetUserById).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.POST("", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.InsertUser).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.POST("/:userId/update", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.UpdateUser).Then(response.SendSuccess).Catch(response.SendError)
	})
}
