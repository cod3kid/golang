package main

import (
	"github.com/cod3kid/golang/10-authentication/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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

	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "No request body found"})
		return
	}

	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(body.Password),10)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash the password"})
		return
	}

	user := models.User{Name: body.Name, Email: body.Email, Password: string(hashedPassword)}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user, "message":"User created"})
}


func main() {
	router := gin.Default()

	models.ConnectDatabase()
	router.GET("/", HelloWorld)
	router.POST("/signup", SignUp)	

	router.Run()
}
