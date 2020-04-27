package utils

import (
	"github.com/barrydev/api-3h-shop/src/constants"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	_corsConfig := cors.DefaultConfig()

	allowOrigins := []string{"http://localhost:3000"}
	allowOrigins = append(allowOrigins, "http://" + constants.WEB_HOST)
	allowOrigins = append(allowOrigins, "https://" + constants.WEB_HOST)

	_corsConfig.AllowOrigins = allowOrigins
	_corsConfig.AllowCredentials = true
	_cors := cors.New(_corsConfig)

	return _cors
}


