package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RequiresAdmin(c *gin.Context) {

	fmt.Println("Requires Admin Middleware triggered")
	c.Next()
}

func RequiresSession(c *gin.Context) {
	token, err := c.Cookie("session_token")
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid session"})
		return
	}
	c.Set("session_token", token)
	c.Next()
}
