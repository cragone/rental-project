package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"server/invoice"

	"github.com/gin-gonic/gin"
)

// invoice will be created with default due date a week from today
func HandleManualInvoice(c *gin.Context) {

	type payload struct {
		Rate      int `json:"rate"`
		TennantID int `json:"tennantID"`
	}

	var p payload

	err := c.BindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var tennant = invoice.Tennant{
		TennantID: p.TennantID,
		Rate:      p.Rate,
	}

	tennants := append([]invoice.Tennant{}, tennant)

	invoice.GenerateInvoices(tennants)

	fmt.Println(tennant)

}

func HandleCreateOrder(c *gin.Context) {
	// get invoice ID
	// check the status
	// if status is not complete initiate

	// get session user from this
	var payload = struct {
		InvoiceID int `json:"invoiceID"`
	}{}

	c.BindJSON(&payload)

	dbname := os.Getenv("POSTGES_DB")
	dbuser := os.Getenv("POSTGRES_USERNAME")
	dbpassword := os.Getenv("POSTGRES_PASSWORD")
	dbhost := os.Getenv("POSTGRES_HOST")

	connString := fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=require port=5432", dbname, dbuser, dbpassword, dbhost)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var status string
	var amount int
	err = db.QueryRow(`SELECT payment_status, amount FROM invoice WHERE payment_id = $1`, payload.InvoiceID).Scan(&status, &amount)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if status == "COMPLETED" {
		c.JSON(400, gin.H{"error": "invoice has already been filled"})
		return
	}

	orderID, err := invoice.GeneratePaypalOrder(amount, payload.InvoiceID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"response": orderID})

}

// This route is accessed by the page paypal redirects to, confirming the order
// need to confirm venmo works the same way
// https://developer.paypal.com/docs/api/webhooks/v1/
//
// Why tonot do this:
// https://stackoverflow.com/questions/36221146/paypal-rest-api-fulfill-order-payment-on-redirect-url-or-on-webhook-call
func HandleConfirmOrder(c *gin.Context) {
	invoiceID := c.Param("id")

	fmt.Println(invoiceID)
	c.JSON(200, gin.H{"response": "done"})
}

func HandleNewOrderTest(c *gin.Context) {

	orderID, err := invoice.GeneratePaypalOrder(320, 500000)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"response": orderID})

}

func HandleOrderStatus(c *gin.Context) {

	// Need to update db to the API response,
	// if API response fails still check db for value
	// if both fail then bad request

	var payload = struct {
		Status string `json:"status"`
	}{}

	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	status, err := invoice.CheckPaypalOrder(payload.Status)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, status)
}
