package main

import (
	"github.com/cod3kid/golang/10-authentication/models"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
}

func SignUp(c *gin.Context) {
	var body struct {
		Name string
		Email string
		Password string
	}

	if c.Bind(&body) == nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "No request body found"})
	}

	fmt.Printf("%+v",body)
	c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
}


func main() {
	router := gin.Default()

	models.ConnectDatabase()
	router.GET("/", HelloWorld)
	router.POST("/signup", SignUp)	

	router.Run()
}
