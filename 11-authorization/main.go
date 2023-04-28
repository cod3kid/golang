package main

import (
	"github.com/cod3kid/golang/11-authorization/models"
	"github.com/cod3kid/golang/11-authorization/middlewares"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
	 "github.com/golang-jwt/jwt/v5"
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
	result := models.DB.Create(&user)

	if result.Error != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create a user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user, "message":"User created"})
}

func Login(c *gin.Context) {
	var body struct {
		Email string
		Password string
	}

	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "No request body found"})
		return
	}

	var user models.User
	models.DB.Where("email = ?", body.Email).First(&user)

	if user.ID == 0{
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(body.Password),10)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash the password"})
		return
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(body.Password)) 
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords doesn't match"})
		return
	}

	// Generate Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"user": user.Email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte("FakeSecret"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenString, "message":"Login Successful"})
}

func GetPosts(c *gin.Context) {
	user,_ := c.Get("user")
	fmt.Println("posts ",user)
	c.JSON(http.StatusOK, gin.H{ "message":"From GetPosts", "data":user})
}



func main() {
	router := gin.Default()

	models.ConnectDatabase()
	router.GET("/", HelloWorld)
	router.POST("/signup", SignUp)
	router.POST("/login", Login)	
	router.GET("/posts",middlewares.RequireAuth,GetPosts)
	

	router.Run()
}
