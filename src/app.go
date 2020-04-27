package src

import (
	"github.com/barrydev/api-3h-shop/src/common/utils"
	"github.com/gin-gonic/gin"
)

type App struct {
	instance *gin.Engine
}

func (app *App) NewGinEngine() *gin.Engine {
	_app := gin.Default()
	_cors := utils.Cors()
	_app.Use(_cors)
	BindRouterWithApp(_app, []gin.HandlerFunc{_cors})

	app.instance = _app

	return _app
}
