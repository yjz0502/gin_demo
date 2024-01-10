package main

import "github.com/gin-gonic/gin"

func main() {
	ginServer := gin.Default()
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"Message": "helloword"})
	})
	ginServer.Run(":8080")
}
