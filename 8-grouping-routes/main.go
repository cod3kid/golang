package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message":"This is getBooks function"})
}

func getBookById(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message":"This is getBookById function"})
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message":"This is getUsers function"})
}

func getUserById(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message":"This is getUserById function"})
}

func main() {
	//router setup
	router := gin.Default()

	// route group for books
	v1 := router.Group("/book")
	{
		v1.GET("/get-books", getBooks)
		v1.GET("/get-book-by-id", getBookById)
	}

	// route group for users
	v2 := router.Group("/user")
	{
		v2.GET("/get-users", getUsers)
		v2.GET("/get-user-by-id", getUserById)
	}

	// Listening to port..  by default it listens to 8080 router.Run()
	router.Run(":8000")
}
