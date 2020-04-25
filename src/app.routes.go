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

	customerRouter := router.Group("/customers")

	routers.BindCustomer(customerRouter)

	/**
	 * Products.
	 */

	productRouter := router.Group("/products")

	routers.BindProduct(productRouter)

	router.POST("/bulk/product/insert", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.BulkInsertProduct).Then(response.SendSuccess).Catch(response.SendError)
	})

	router.POST("/bulk/product/update", func(c *gin.Context) {
		handle := response.Handle{Context: c}

		handle.Try(controllers.BulkUpdateProduct).Then(response.SendSuccess).Catch(response.SendError)
	})

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
