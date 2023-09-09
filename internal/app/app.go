package app

import (
	"encryption-algorithm/pkg/logger"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	logger.InitZap()

	r := InitGinEngine()

	return r
}
