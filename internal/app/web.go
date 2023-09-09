package app

import (
	"encryption-algorithm/internal/app/router"
	"github.com/gin-gonic/gin"
)

func InitGinEngine() *gin.Engine {
	app := gin.Default()

	router.Register(app)

	return app
}
