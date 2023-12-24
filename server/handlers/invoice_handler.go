package handlers

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"server/invoice"

	"github.com/gin-gonic/gin"
	"github.com/plutov/paypal/v4"
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
	type Payload struct {
		InvoiceID string `json:"invoiceID"`
	}

	var payload Payload

	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

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
		fmt.Println("test")
		fmt.Println(payload.InvoiceID)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if status == "COMPLETE" {
		c.JSON(400, gin.H{"error": "invoice has already been filled"})
		return
	}

	orderID, err := invoice.GeneratePaypalOrder(amount, payload.InvoiceID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err = db.Exec(`INSERT INTO invoice_paypal_lookup (invoice_id, paypal_order_id) VALUES ($1, $2)`, payload.InvoiceID, orderID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"response": orderID})

}

func HandleTakeWebhookResponse(c *gin.Context) {

	// read the json body into a string
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("Error reading request body")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading request body"})
		return
	}

	// This is the only situation where we would potentially send a 400
	// verify the webhook signiature
	err = verifyWebhookSignature(c, body)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Roll back the body read
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	var payload struct {
		Resource struct {
			ID string `json:"id"`
		} `json:"resource"`
	}

	err = c.BindJSON(&payload)
	if err != nil {
		c.JSON(200, gin.H{"error": "issue parsing request"})
		return
	}

	fmt.Println(payload.Resource.ID)

	dbname := os.Getenv("POSTGES_DB")
	dbuser := os.Getenv("POSTGRES_USERNAME")
	dbpassword := os.Getenv("POSTGRES_PASSWORD")
	dbhost := os.Getenv("POSTGRES_HOST")

	connString := fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=require port=5432", dbname, dbuser, dbpassword, dbhost)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	var invoiceID string

	err = db.QueryRow(`SELECT invoice_id FROM invoice_paypal_lookup WHERE paypal_order_id = $1`, payload.Resource.ID).Scan(&invoiceID)
	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}

	_, err = db.Exec(`UPDATE invoice SET payment_status = 'COMPLETE' WHERE payment_id = $1`, invoiceID)
	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"response": payload.Resource.ID})
}

// This verifies the webhook with the server to stop bad actors.
// Must have correct PAYPAL_WEBHOOK_ID corresponding to the webook
func verifyWebhookSignature(c *gin.Context, body []byte) error {
	// Get the signature-related headers from the request headers
	transmissionID := c.GetHeader("Paypal-Transmission-Id")
	transmissionSignature := c.GetHeader("Paypal-Transmission-Sig")
	transmissionTime := c.GetHeader("Paypal-Transmission-Time")
	certURL := c.GetHeader("Paypal-Cert-Url")

	verifyPayload := fmt.Sprintf(`{
		"transmission_id": "%s",
		"transmission_time": "%s",
		"cert_url": "%s",
		"auth_algo": "SHA256withRSA",
		"transmission_sig": "%s",
		"webhook_id": "%s",
		"webhook_event": %s
	  }`, transmissionID, transmissionTime, certURL, transmissionSignature, os.Getenv("PAYPAL_WEBHOOK_ID"), string(body))

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/v1/notifications/verify-webhook-signature", os.Getenv("PAYPAL_BASE_URL")),
		bytes.NewBufferString(verifyPayload),
	)
	if err != nil {
		return fmt.Errorf("Error creating HTTP request: %v", err)
	}

	clientID := os.Getenv("PAYPAL_CLIENT_ID")
	secret := os.Getenv("PAYPAL_SECRET")

	apiURL := os.Getenv("PAYPAL_BASE_URL")

	// Create a client instance
	client, err := paypal.NewClient(clientID, secret, apiURL)
	if err != nil {
		return fmt.Errorf("Error creating HTTP request: %v", err)
	}
	client.SetLog(os.Stdout) // Set log to terminal stdout

	accessToken, err := client.GetAccessToken(context.Background())
	if err != nil {
		return fmt.Errorf("Error creating HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken.Token)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	verificationStatus := struct {
		Status string `json:"verification_status"`
	}{}

	err = json.Unmarshal(buf, &verificationStatus)
	if err != nil {
		return err
	}

	fmt.Println(verificationStatus.Status)
	if verificationStatus.Status != "SUCCESS" {
		return errors.New("could not confirm the webhook")
	}

	return nil
}

// This route is accessed by the page paypal redirects to, confirming the order
// need to confirm venmo works the same way
// https://developer.paypal.com/docs/api/webhooks/v1/
//
// Why to not do this:
// https://stackoverflow.com/questions/36221146/paypal-rest-api-fulfill-order-payment-on-redirect-url-or-on-webhook-call
func HandleConfirmOrder(c *gin.Context) {
	invoiceID := c.Param("id")

	fmt.Println(invoiceID)
	c.JSON(200, gin.H{"response": "done"})
}

func HandleNewOrderTest(c *gin.Context) {

	orderID, err := invoice.GeneratePaypalOrder(320, "sdnkfnseongosennke")
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
