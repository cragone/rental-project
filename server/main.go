package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hello, World!")
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r := gin.Default()

	auth := r.Group("/auth")
	auth.GET("/newuser", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{"response": "good"})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{"response": "default"})
	})

	r.Run(":80")
}
