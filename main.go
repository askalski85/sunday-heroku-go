package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
  router.GET("/foo", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "bar",
		})
	})
	router.Run(":8085")
}
