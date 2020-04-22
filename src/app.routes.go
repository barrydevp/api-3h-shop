package src

import (
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/controllers"
	"github.com/barrydev/api-3h-shop/src/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BindRouterWithApp(router *gin.Engine, handlerFuncs []gin.HandlerFunc) {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "API 3H-Shop.")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "API 3H-Shop:pong")
	})

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

	customerRouter := router.Group("/customers", handlerFuncs...)

	routers.BindCustomer(customerRouter)

	/**
	 * Products.
	 */

	productRouter := router.Group("/products", handlerFuncs...)

	routers.BindProduct(productRouter)

	//router.POST("/bulk/product", func(c *gin.Context) {
	//	handle := response.Handle{Context: c}
	//
	//	handle.Try(controllers.BulkInsertProduct).Then(response.SendSuccess).Catch(response.SendError)
	//})

	/**
	 * ProductItems.
	 */

	productItemRouter := router.Group("/product-items", handlerFuncs...)

	routers.BindProductItem(productItemRouter)

	/**
	 * Orders.
	 */

	orderRouter := router.Group("/orders", handlerFuncs...)

	routers.BindOrder(orderRouter)

	/**
	 * OrderItems.
	 */

	orderItemRouter := router.Group("/order-items", handlerFuncs...)

	routers.BindOrderItem(orderItemRouter)


	/**
	 * Shippings.
	 */

	shippingRouter := router.Group("/shippings", handlerFuncs...)

	routers.BindShipping(shippingRouter)


	/**
	 * Current.
	 */

	currentRouter := router.Group("/current", handlerFuncs...)

	routers.BindCurrent(currentRouter)
}
