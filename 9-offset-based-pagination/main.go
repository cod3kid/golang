package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message":"This is getBooks function"})
}

func main() {
	//router setup
	router := gin.Default()

	// route group for books
	router.GET("/get-books", getBooks)

	// Listening to port..  by default it listens to 8080 router.Run()
	router.Run(":8000")
}
