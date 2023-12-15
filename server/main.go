package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"server/handlers"
	"server/invoice"
	"server/middleware"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"*"}

	r := gin.Default()

	r.SetTrustedProxies(nil)

	// Configure header controls middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	r.LoadHTMLGlob("templates/*")

	r.Static("/assets", "./assets")

	auth := r.Group("/auth")
	auth.POST("/login", func(c *gin.Context) { fmt.Println("login route triggered") })

	admin := r.Group("/admin")
	admin.Use(middleware.RequiresAdmin)

	admin.GET("/registeruser", handlers.RegisterUser)

	property := r.Group("property")
	{
		property.POST("/create", CreateProperty)
		property.POST("/delete", DeleteProperty)
		property.GET("/list", PropertyList)
	}

	tennant := r.Group("/tennant")
	{
		tennant.POST("/create", handlers.CreateTennant)
		tennant.GET("/get", handlers.GetTennant)
		tennant.GET("/property_tennants", PropertyTennantIDList)
	}

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// CRON job for setting up invoices
	s := gocron.NewScheduler(time.UTC)
	// How often do we run the task?
	// In prod
	s.Every(1).Day().At("2:00").Do(invoice.GenerateAllInvoices)
	// s.Every(100).Second().Do(invoice.GenerateAllInvoices)
	// start the scheduler
	s.StartAsync()

	fmt.Println("Server Started")

	id, err := invoice.GeneratePaypalOrder(300)
	if err != nil {
		panic(err)
	}
	fmt.Println("order id:")
	fmt.Println(id)

	x, err := invoice.CheckPaypalOrder("2AR53697HY170130S")
	if err != nil {
		panic(err)
	}

	fmt.Println("order status:")
	fmt.Println(x)

	order := r.Group("/order")
	{
		order.GET("/new", handlers.HandleNewOrder)
		order.POST("/status", handlers.HandleOrderStatus)
	}

	r.Run(":80")
}

type Property struct {
	Address string `json:"address"`
	ID      int    `json:"id"`
}

func CreateProperty(c *gin.Context) {

	var currentProperty Property

	err := c.ShouldBindJSON(&currentProperty)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(currentProperty.Address)

	dbname := os.Getenv("POSTGES_DB")
	dbuser := os.Getenv("POSTGRES_USERNAME")
	dbpassword := os.Getenv("POSTGRES_PASSWORD")
	dbhost := os.Getenv("POSTGRES_HOST")

	connString := fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=require port=5432", dbname, dbuser, dbpassword, dbhost)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	_, err = db.Query("INSERT INTO property_info (property_address) VALUES ($1)", currentProperty.Address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "successfully added property"})
}

func DeleteProperty(c *gin.Context) {
	var currentProperty Property

	err := c.ShouldBindJSON(&currentProperty)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(currentProperty.Address)

	dbname := os.Getenv("POSTGES_DB")
	dbuser := os.Getenv("POSTGRES_USERNAME")
	dbpassword := os.Getenv("POSTGRES_PASSWORD")
	dbhost := os.Getenv("POSTGRES_HOST")

	connString := fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=require port=5432", dbname, dbuser, dbpassword, dbhost)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	res, err := db.Exec("DELETE FROM property_info WHERE property_id = $1", currentProperty.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	rows, err := res.RowsAffected()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	defer db.Close()

	if rows == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nothing to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "successfully deleted property"})
}

func PropertyList(c *gin.Context) {
	dbname := os.Getenv("POSTGES_DB")
	dbuser := os.Getenv("POSTGRES_USERNAME")
	dbpassword := os.Getenv("POSTGRES_PASSWORD")
	dbhost := os.Getenv("POSTGRES_HOST")

	connString := fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=require port=5432", dbname, dbuser, dbpassword, dbhost)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT property_id, property_address FROM property_info")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	var properties []Property

	for rows.Next() {
		var property Property
		err = rows.Scan(&property.ID, &property.Address)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		properties = append(properties, property)
	}

	c.JSON(http.StatusOK, gin.H{"propertyList": properties})
}

func PropertyTennantIDList(c *gin.Context) {
	var payload = struct {
		PropertyID int `json:"propertyID"`
	}{}

	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbname := os.Getenv("POSTGES_DB")
	dbuser := os.Getenv("POSTGRES_USERNAME")
	dbpassword := os.Getenv("POSTGRES_PASSWORD")
	dbhost := os.Getenv("POSTGRES_HOST")

	connString := fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=require port=5432", dbname, dbuser, dbpassword, dbhost)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT b_id FROM brokie WHERE property_id = $1", payload.PropertyID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	var tennantIDs []int

	for rows.Next() {
		var tennantID int
		err = rows.Scan(&tennantID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		tennantIDs = append(tennantIDs, tennantID)

	}

	c.JSON(http.StatusBadRequest, gin.H{"response": tennantIDs})

}
