package router

import (
	"encryption-algorithm/internal/app/api"
	"github.com/gin-gonic/gin"
)

func Register(app *gin.Engine) error {
	RegisterAPI(app)

	return nil
}

func RegisterAPI(app *gin.Engine) {
	app.Group("/test")
	{
		app.GET("/ping", api.Ping)
	}
}
