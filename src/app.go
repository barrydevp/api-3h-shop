package src

import (
	"github.com/barrydev/api-3h-shop/src/constants"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
	instance *gin.Engine
}

func (app *App) NewGinEngine() *gin.Engine {
	_app := gin.Default()
	_corsConfig := cors.DefaultConfig()
	_corsConfig.AllowOrigins = []string{"localhost:3000", constants.PRIMARY_HOST}
	_corsConfig.AllowCredentials = true
	_cors := cors.New(_corsConfig)
	_app.Use(_cors)
	BindRouterWithApp(_app, []gin.HandlerFunc{_cors})

	app.instance = _app

	return _app
}
