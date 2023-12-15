package invoice

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/plutov/paypal/v4"
)

func GenerateAllInvoices() {
	tennants := pullValidTennants()
	GenerateInvoices(tennants)

}

type Tennant struct {
	TennantID int
	Rate      int
}

func pullValidTennants() []Tennant {

	// generate an invoice a week before it is due
	day := time.Now().Add(time.Hour * 24 * 7).Day()

	// Criteria:
	// Must be the correct day of month
	// Must be active

	// *** Below is not handled in the DB architecture
	// Todays date cannot be after the last payment day

	dbname := os.Getenv("POSTGES_DB")
	dbuser := os.Getenv("POSTGRES_USERNAME")
	dbpassword := os.Getenv("POSTGRES_PASSWORD")
	dbhost := os.Getenv("POSTGRES_HOST")

	connString := fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=require port=5432", dbname, dbuser, dbpassword, dbhost)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT b_id, rent_rate FROM brokie WHERE payment_day = $1 AND active = 'Y'", day)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var validTennants []Tennant

	for rows.Next() {
		var ten Tennant
		err = rows.Scan(&ten.TennantID, &ten.Rate)
		if err != nil {
			panic(err)
		}
		validTennants = append(validTennants, ten)
	}

	return validTennants
}

func GenerateInvoices(tennants []Tennant) {

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

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	for _, tennant := range tennants {
		_, err = tx.Exec("INSERT INTO invoice (due_date, tennant_id, amount) VALUES ((CURRENT_DATE - INTERVAL '7 days'), $1, $2)", tennant.TennantID, tennant.Rate)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	// TODO mail a report to admin of the invoices generated
	fmt.Println("sucessfully generated invoices")
}

type TokenResponse struct {
	Scope       string `json:"scope"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	AppID       string `json:"app_id"`
	ExpiresIn   int    `json:"expires_in"`
	Nonce       string `json:"nonce"`
}

func GeneratePaypalOrder(amount int) (string, error) {

	// Prepare the request body using environment variables
	// orderIntent := "CAPTURE"
	clientID := os.Getenv("PAYPAL_CLIENT_ID")
	secret := os.Getenv("PAYPAL_SECRET")

	apiURL := os.Getenv("PAYPAL_BASE_URL")

	// Create a client instance
	c, err := paypal.NewClient(clientID, secret, apiURL)
	if err != nil {
		return "", err
	}

	accessToken, err := c.GetAccessToken(context.Background())
	if err != nil {
		return "", err
	}

	rawOrderJSON := fmt.Sprintf(`{
		"intent": "CAPTURE",
		"purchase_units": [
			{
				"amount": {
					"currency_code": "USD",
					"value": "%d.00"
				  }
			}
		]
		
	}`, amount)

	// Create the HTTP request for order creation
	req, err := http.NewRequest("POST", apiURL+"/v2/checkout/orders", bytes.NewBuffer([]byte(rawOrderJSON)))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken.Token)

	// Send the order creation request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read and print the response body
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	type body struct {
		ID string `json:"id"`
	}

	var orderID body

	err = json.Unmarshal(buf, &orderID)
	if err != nil {
		return "", err
	}

	return orderID.ID, nil
}

func CheckPaypalOrder(orderID string) (string, error) {
	// sample 143534824A8662459
	clientID := os.Getenv("PAYPAL_CLIENT_ID")
	secret := os.Getenv("PAYPAL_SECRET")

	apiURL := os.Getenv("PAYPAL_BASE_URL")

	// Create a client instance
	c, err := paypal.NewClient(clientID, secret, apiURL)
	if err != nil {
		return "", err
	}
	c.SetLog(os.Stdout) // Set log to terminal stdout

	accessToken, err := c.GetAccessToken(context.Background())
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("GET", apiURL+fmt.Sprintf("/v2/checkout/orders/%s", orderID), nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken.Token)

	// Send the order creation request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	type body struct {
		Status string `json:"status"`
	}

	var status body

	err = json.Unmarshal(buf, &status)
	if err != nil {
		return "", err
	}

	return status.Status, nil
}
