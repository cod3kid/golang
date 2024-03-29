package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//router setup
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	//routes
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	// Listening to port..  by default it listens to 8080 router.Run()
	router.Run(":8000")
}
