package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Hello World Env"})
}

func main() {
	router := gin.Default()

	router.GET("/", HelloWorld)

	router.Run()
}
