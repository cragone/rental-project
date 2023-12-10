package invoice

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

func GenerateAllInvoices() {
	fmt.Println("all invoices have been generated")
	pullValidTennants()
}

type Tennant struct {
	TennantID int
	Rate      int
}

func pullValidTennants() {

	// Get the day it is in the month
	day := time.Now().Day()

	fmt.Println(day)

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

	rows, err := db.Query("SELECT b_id, amount FROM brokie WHERE payment_day = $1 AND active = 'Y'")
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

	// Not done need to test and pass data
	fmt.Println(validTennants)

}
