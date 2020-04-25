package src

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
	instance *gin.Engine
}

func (app *App) NewGinEngine() *gin.Engine {
	_app := gin.Default()
	_corsConfig := cors.DefaultConfig()
	_corsConfig.AllowAllOrigins = true
	_corsConfig.AllowCredentials = true
	_cors := cors.New(_corsConfig)
	_app.Use(_cors)
	BindRouterWithApp(_app, []gin.HandlerFunc{_cors})

	app.instance = _app

	return _app
}
