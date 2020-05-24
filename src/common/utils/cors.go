package utils

import (
	"github.com/barrydev/api-3h-shop/src/constants"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func _generateAndAppendURL(listURL *[]string, uri string) {
	*listURL = append(*listURL, "http://"+uri)
	*listURL = append(*listURL, "http://www."+uri)
	*listURL = append(*listURL, "https://"+uri)
	*listURL = append(*listURL, "https://www."+uri)
}

func Cors() gin.HandlerFunc {
	_corsConfig := cors.DefaultConfig()

	allowOrigins := []string{}
	_generateAndAppendURL(&allowOrigins, "localhost:3000")
	_generateAndAppendURL(&allowOrigins, constants.WEB_HOST)
	_generateAndAppendURL(&allowOrigins, "test-cors.org")
	// allowOrigins = append(allowOrigins, "http://"+constants.WEB_HOST)
	// allowOrigins = append(allowOrigins, "https://.test-cors.org")
	// allowOrigins = append(allowOrigins, "http://.test-cors.org")
	// allowOrigins = append(allowOrigins, "https://www.test-cors.org")
	// allowOrigins = append(allowOrigins, "http://www.test-cors.org")
	// allowOrigins = append(allowOrigins, "https://"+constants.WEB_HOST)

	allowHeaders := []string{"Origin", "Authorization"}

	_corsConfig.AllowOrigins = allowOrigins
	_corsConfig.AllowCredentials = true
	_corsConfig.AllowHeaders = allowHeaders
	_corsConfig.AllowFiles = true
	_cors := cors.New(_corsConfig)

	return _cors
}
