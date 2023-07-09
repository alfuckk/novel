package oss

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func RegisterHandlers(engine *gin.Engine) {
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})
}