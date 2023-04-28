package middlewares

import (	
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/cod3kid/golang/11-authorization/models"
	"fmt"
	"time"
	"net/http"
)
func RequireAuth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
fmt.Println("Inside the Middleware",tokenString)

token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	return []byte("FakeSecret"), nil
})

if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	fmt.Println(claims["user"])

	if float64(time.Now().Unix()) > claims["exp"].(float64){
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	var user models.User
	models.DB.Where("email = ?", claims["user"]).First(&user)

	if user.ID == 0{
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	c.Set("user",user)
	c.Next()
} else {
	fmt.Println(err)
	c.AbortWithStatus(http.StatusUnauthorized)
}
}