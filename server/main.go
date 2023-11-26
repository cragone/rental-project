package main

import (
	"fmt"
	"net/http"
	"server/handlers"
	"server/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.Static("/assets", "./assets")

	auth := r.Group("/auth")
	auth.POST("/login", func(c *gin.Context) { fmt.Println("login route triggered") })

	admin := r.Group("/admin")
	admin.Use(middleware.RequiresAdmin)

	admin.GET("/registeruser", handlers.RegisterUser)

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	fmt.Println("Server Started")
	fmt.Println("here is my change")
	r.Run(":80")
}
