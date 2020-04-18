package src

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	instance *gin.Engine
}

func (app *App) NewGinEngine() *gin.Engine {
	_app := gin.Default()

	//_app.Use(cors.Default())
	BindRouterWithApp(_app)

	app.instance = _app

	return _app
}
