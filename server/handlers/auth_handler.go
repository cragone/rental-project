package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	fmt.Println("hello")
}

// this is an admin only route
func ProvisionAccount(c *gin.Context) {
	// POST req type

	var payload = struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}{}

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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	res, err := db.Exec(`INSERT INTO user_info (user_type, full_name, email, one_time_code) VALUES ('user', $1, $2, 'google_auth')`, payload.Name, payload.Email)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	n, err := res.RowsAffected()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if n == 0 {
		c.JSON(400, gin.H{"error": "user already exists in system with that email"})
		return
	}

	c.JSON(200, gin.H{"response": "create new user in system"})

}

func GenerateGoogleAuthURI(c *gin.Context) {

	clientID := os.Getenv("CLIENT_ID")
	redirectURI := os.Getenv("REDIRECT_URI")

	googleURI := fmt.Sprintf("https://accounts.google.com/o/oauth2/auth?client_id=%s&redirect_uri=%s&response_type=code&scope=openid+profile+email&prompt=consent",
		clientID, redirectURI)

	// Gin auto escapes & valuesto unicode for some reason so we must use this bypass
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	fmt.Fprintf(c.Writer, `{"response":"%s"}`, googleURI)
}

func ValidateGoogleHandshake(c *gin.Context) {
	var payload = struct {
		Code string `json:"code"`
	}{}

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	redirectURI := os.Getenv("REDIRECT_URI")

	toSend := fmt.Sprintf(`{
	"code": "%s",
	"client_id": "%s",
	"client_secret": "%s",
	"redirect_uri": "%s",
	"grant_type": "authorization_code"
	}`, payload.Code, clientID, clientSecret, redirectURI)

	req, err := http.NewRequest("POST", "https://oauth2.googleapis.com/token", bytes.NewBuffer([]byte(toSend)))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the order creation request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	// Read and print the response body
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	type body struct {
		ID string `json:"access_token"`
	}

	var accessToken body

	err = json.Unmarshal(buf, &accessToken)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"response": accessToken})

}
