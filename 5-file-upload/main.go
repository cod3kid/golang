package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message":"Hello"})
}

func main() {
	//router setup
	router := gin.Default()
	router.Static("/assets", "./assets")

	//routes
	router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, "uploads/"+file.Filename)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	// Listening to port..  by default it listens to 8080 router.Run()
	router.Run(":8000")
}
