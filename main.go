package main

import (
	"go-mysql/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		infrastructure.NewDatabase()
		context.JSON(http.StatusOK, gin.H{"data": "Hello World !"})
	})
	router.Run(":8000")
}
