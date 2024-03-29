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
	{
		auth.POST("/login", func(c *gin.Context) { fmt.Println("login route triggered") })
		auth.GET("/google_uri", handlers.GenerateGoogleAuthURI)
		auth.POST("/google_session_handshake", handlers.ValidateGoogleHandshake)
		auth.GET("/session", handlers.GetSession)
	}

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
		tennant.POST("/get", handlers.GetTennant)
		tennant.POST("/property_tennants", PropertyTennantIDList)
	}

	payment := r.Group("/payment")
	{
		payment.GET("/list_user", listUserPayments)
	}

	// CRON job for setting up invoices
	s := gocron.NewScheduler(time.UTC)
	// How often do we run the task?
	// In prod 1 time per day at 2am
	// 2am in EST is 7am UTC
	s.Every(1).Day().At("7:00").Do(invoice.GenerateAllInvoices)
	// start the scheduler
	s.StartAsync()

	fmt.Println("Server Started")

	order := r.Group("/invoice")
	{
		// WHAT WE HAVE RIGHT NOW WILL NOT WORK!!
		// orders expire so we need a new strategy.
		// We cant save the order ID so we need to send the invoice ID within the redirect link to track the payment.
		// Only when the redirect link is triggered we can track the invoice
		order.GET("/test_new", handlers.HandleNewOrderTest)

		order.POST("/manual_invoice", handlers.HandleManualInvoice)
		order.POST("/create_order", handlers.HandleCreateOrder)

		order.POST("/webhook", handlers.HandleTakeWebhookResponse)
	}

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run(":80")
}

func listUserPayments(c *gin.Context) {
	output := `
	[
    { 'id': 1, 'due': 100, 'due_date': '2023-12-01', 'type': 'Utility' },
    { 'id': 2, 'due': 200, 'due_date': '2023-12-05', 'type': 'Rent' },
    { 'id': 3, 'due': 150, 'due_date': '2023-12-10', 'type': 'Late Fee' }
	]
	`
	c.JSON(200, output)
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

	c.JSON(200, gin.H{"response": tennantIDs})

}
