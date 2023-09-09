package api

import (
	"encryption-algorithm/pkg/error"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	error.Response(c, error.OK, gin.H{}, "pong")
}
