package middlewares

import (
	"net/http"

	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/gin-gonic/gin"
)

func InjectPayloadToContext(c *gin.Context, claims *factories.AccessTokenClaims) {
	c.Set("payload_token", claims)
}

func AuthenticateJwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := factories.RetriveAccessTokenPayload(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		InjectPayloadToContext(c, claims)

		c.Next()
	}
}
