package main

import (
	"github.com/cod3kid/golang/9-offset-based-pagination/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
	var pokemons []models.Pokemon
	models.DB.Limit(2).Offset(2).Find(&pokemons)
	c.JSON(http.StatusOK, gin.H{"data": pokemons})
}


func main() {
	router := gin.Default()

	models.ConnectDatabase()
	router.GET("/read-all", FindBooks)
	

	router.Run()
}
