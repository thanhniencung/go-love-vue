package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
