package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	//router setup
	router := gin.Default()

	// Listening to port..  by default it listens to 8080 router.Run()
	router.Run(":8000")
}
