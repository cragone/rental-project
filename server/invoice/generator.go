package invoice

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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

func GeneratePaypal() {

	client, err := paypal.NewClient(os.Getenv("PAYPAL_CLIENT_ID"), os.Getenv("PAYPAL_SECRET"), "https://api.sandbox.paypal.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(client)

	ctx := context.Background()

	order, err := client.CreateOrder(ctx, "CAPTURE",
		[]paypal.PurchaseUnitRequest{
			{
				Amount: &paypal.PurchaseUnitAmount{
					Currency: "USD",
					Value:    "20.00",
				},
			},
		}, &paypal.CreateOrderPayer{}, &paypal.ApplicationContext{})

	if err != nil {
		panic(err)
	}
	fmt.Println("below here")
	fmt.Println(order.ID)

}
