package main

import (
	"os"
  "github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	router := gin.Default()
  router.GET("/foo", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "bar",
		})
	})
	router.Run(":" + port)
}
