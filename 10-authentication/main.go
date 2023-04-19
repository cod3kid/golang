package main

import (
	"github.com/cod3kid/golang/10-authentication/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
}


func main() {
	router := gin.Default()

	models.ConnectDatabase()
	router.GET("/", HelloWorld)
	

	router.Run()
}
