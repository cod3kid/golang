package middlewares

import (	
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"fmt"
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
	c.Set("user",claims["user"])
	c.Next()
} else {
	fmt.Println(err)
	c.AbortWithStatus(http.StatusUnauthorized)
}
}