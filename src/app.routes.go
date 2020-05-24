package src

import (
	"net/http"

	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/controllers"
	"github.com/barrydev/api-3h-shop/src/middlewares"
	"github.com/barrydev/api-3h-shop/src/routers"
	"github.com/gin-gonic/gin"
)

func BindRouterWithApp(router *gin.Engine, handlerFuncs []gin.HandlerFunc) {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "API 3H-Shop.")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "API 3H-Shop:pong")
	})

	router.POST("/admin/authenticate", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.AuthenticateAdmin).Then(response.SendSuccess).Catch(response.SendError)
	})

	/**
	 * Router.
	 */
	BindRouter(router, handlerFuncs)

	/**
	 * AuthRouter.
	 */
	authRouter := router.Group("/auth")
	authRouter.Use(middlewares.AuthenticateJwtToken())
	BindAuthRouter(authRouter, handlerFuncs)

	/**
	 * AdminRouter.
	 */
	adminRouter := router.Group("/admin")
	adminRouter.Use(middlewares.AuthenticateAdminJwtToken())
	BindAdminRouter(adminRouter, handlerFuncs)

	testRouter := router.Group("/test")
	testRouter.Use(middlewares.AuthenticateJwtToken())
	testRouter.GET("/payload", func(c *gin.Context) {
		payload, ok := c.Get("payload_token")

		if !ok {
			c.String(400, "invalid payload")

			return
		}

		c.JSON(200, payload)
	})
}

func BindRouter(router *gin.Engine, handlerFuncs []gin.HandlerFunc) {
	/**
	 * Categories.
	 */

	categoryRouter := router.Group("/categories")

	routers.BindCategory(categoryRouter)

	router.GET("/category-tree/all", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.GetAllCategoryTree).Then(response.SendSuccess).Catch(response.SendError)
	})

	/**
	 * Customers.
	 */

	customerRouter := router.Group("/customers")

	routers.BindCustomer(customerRouter)

	/**
	 * Users.
	 */

	userRouter := router.Group("/users")

	routers.BindUser(userRouter)

	/**
	 * Products.
	 */

	productRouter := router.Group("/products")

	routers.BindProduct(productRouter)

	// router.POST("/bulk/product/insert", func(c *gin.Context) {
	// 	handle := response.Handle{Context: c}

	// 	handle.Try(controllers.BulkInsertProduct).Then(response.SendSuccess).Catch(response.SendError)
	// })

	// router.POST("/bulk/product/update", func(c *gin.Context) {
	// 	handle := response.Handle{Context: c}

	// 	handle.Try(controllers.BulkUpdateProduct).Then(response.SendSuccess).Catch(response.SendError)
	// })

	/**
	 * ProductItems.
	 */

	productItemRouter := router.Group("/product-items")

	routers.BindProductItem(productItemRouter)

	/**
	 * Orders.
	 */

	orderRouter := router.Group("/orders")

	routers.BindOrder(orderRouter)

	/**
	 * OrderItems.
	 */

	orderItemRouter := router.Group("/order-items")

	routers.BindOrderItem(orderItemRouter)

	/**
	 * Shippings.
	 */

	shippingRouter := router.Group("/shippings")

	routers.BindShipping(shippingRouter)

	/**
	 * Current.
	 */

	currentRouter := router.Group("/current")

	routers.BindCurrent(currentRouter)

}

func BindAuthRouter(router *gin.RouterGroup, handlerFuncs []gin.HandlerFunc) {
	/**
	 * Customers.
	 */

	// customerRouter := router.Group("/customers")

	// routers.BindCustomer(customerRouter)

	/**
	 * Users.
	 */

	userRouter := router.Group("/users")

	routers.BindUserAuth(userRouter)

	/**
	 * Orders.
	 */

	// orderRouter := router.Group("/orders")

	// routers.BindOrder(orderRouter)

	/**
	 * OrderItems.
	 */

	// orderItemRouter := router.Group("/order-items")

	// routers.BindOrderItem(orderItemRouter)

	/**
	 * Shippings.
	 */

	// shippingRouter := router.Group("/shippings")

	// routers.BindShipping(shippingRouter)

	/**
	 * Current.
	 */

	currentRouter := router.Group("/current")

	routers.BindCurrentAuth(currentRouter)

}

func BindAdminRouter(router *gin.RouterGroup, handlerFuncs []gin.HandlerFunc) {
	/**
	 * Customers.
	 */

	customerRouter := router.Group("/customers")

	routers.BindCustomerAdmin(customerRouter)

	/**
	 * Customers.
	 */

	categoryRouter := router.Group("/categories")

	routers.BindCategoryAdmin(categoryRouter)

	/**
	 * Products.
	 */

	productRouter := router.Group("/products")

	routers.BindProductAdmin(productRouter)

	/**
	 * Users.
	 */

	userRouter := router.Group("/users")

	routers.BindUserAdmin(userRouter)

	/**
	 * Orders.
	 */

	orderRouter := router.Group("/orders")

	routers.BindOrderAdmin(orderRouter)

	/**
	 * OrderItems.
	 */

	// orderItemRouter := router.Group("/order-items")

	// routers.BindOrderItem(orderItemRouter)

	/**
	 * Shippings.
	 */

	shippingRouter := router.Group("/shippings")

	routers.BindShippingAdmin(shippingRouter)

	/**
	 * Current.
	 */

	currentRouter := router.Group("/current")

	routers.BindCurrentAuth(currentRouter)

}
