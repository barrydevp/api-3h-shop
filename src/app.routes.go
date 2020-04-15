package src

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Test struct {
	Hello string `json:"hello"`
	World string `json:"world"`
}

func BindRouterWithApp(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "API 3H-Shop.")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "API 3H-Shop:pong")
	})

	router.GET("/test", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(func(c *gin.Context) (interface{}, error) {
			return Test{
				Hello: "ok",
				World: "ok",
			}, nil
		}).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.GET("/error", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(func(c *gin.Context) (interface{}, error) {
			return nil, errors.New("test error")
		}).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.GET("/users", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetListUser).Then(response.SendSuccess).Catch(response.SendError)
	})
}
