package main

import (
	"golang/2-crud-postgres/controllers"
	"golang/2-crud-postgres/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	books := router.Group("/api/books")
	{

		books.GET("/read-all", controllers.FindBooks)
		books.POST("/create", controllers.CreateBook)
		books.GET("/read/:id", controllers.FindBook)
		books.PUT("/update/:id", controllers.UpdateBook)
		books.DELETE("/delete/:id", controllers.DeleteBook)

	}

	router.Run()
}
