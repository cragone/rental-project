package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Tennant struct {
	PaymentDay   int    `json:"paymentDay"`
	RentRate     int    `json:"rentRate"`
	PropertyID   int    `json:"propertyID"`
	ActiveStatus string `json:"activeStatus"`
	Email        string `json:"email"`
}

func CreateTennant(c *gin.Context) {
	var tennant Tennant

	err := c.BindJSON(&tennant)
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

	fmt.Println(tennant)

	res, err := db.Exec("INSERT INTO brokie (payment_day, rent_rate, property_id, active, email) VALUES ($1, $2, $3, $4, $5)",
		tennant.PaymentDay,
		tennant.RentRate,
		tennant.PropertyID,
		tennant.ActiveStatus,
		tennant.Email)
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

	// This should contact the tennant and tell status or something

	c.JSON(http.StatusAccepted, gin.H{"response": "accepted"})

}

func GetTennant(c *gin.Context) {
	var payload = struct {
		TennantID int `json:"tennantID"`
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

	var tennant Tennant

	query := `SELECT 
		payment_day,
		rent_rate,
		property_id,
		active,
		email 
		FROM brokie WHERE b_id = $1`

	row := db.QueryRow(query, payload.TennantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = row.Scan(&tennant.PaymentDay, &tennant.RentRate, &tennant.PropertyID, &tennant.ActiveStatus, &tennant.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tennant": tennant})

}
