package middlewares

import (	
	"github.com/gin-gonic/gin"
	"fmt"
)
func RequireAuth(c *gin.Context) {
fmt.Println("Inside the Middleware")
	c.Next()
}