package main

import (
	"github.com/gin-gonic/gin"
)
func main() {
	r := gin.Default()
	r.GET("/demo", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	r.Run(":8080") // Start the server on port 8080
}