package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RequiresAdmin(c *gin.Context) {

	fmt.Println("Requires Admin Middleware triggered")
	c.Next()
}
