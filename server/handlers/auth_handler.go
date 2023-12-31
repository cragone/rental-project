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
	"github.com/google/uuid"
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

	err = CreateAccount(payload.Email, payload.Name)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"response": "create new user in system"})

}

func CreateAccount(email string, name string) error {
	dbname := os.Getenv("POSTGES_DB")
	dbuser := os.Getenv("POSTGRES_USERNAME")
	dbpassword := os.Getenv("POSTGRES_PASSWORD")
	dbhost := os.Getenv("POSTGRES_HOST")

	connString := fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=require port=5432", dbname, dbuser, dbpassword, dbhost)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return err
	}
	defer db.Close()

	res, err := db.Exec(`INSERT INTO user_info (user_type, full_name, email, one_time_code) VALUES ('user', $1, $2, 'google_auth')`, name, email)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return err
	}

	return nil
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

	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

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

	var response = struct {
		AccessToken string `json:"access_token"`
	}{}

	err = json.Unmarshal(buf, &response)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	email, name, err := getGoogleInfo(response.AccessToken)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	sessionToken, err := CreateSession(email)
	if err != nil {
		c.JSON(400, gin.H{"error": "contact administrator for account"})
		return
	}
	fmt.Println(sessionToken)

	c.SetCookie("session_token", sessionToken, 100000, "/", "localhost", false, true)

	c.JSON(200, gin.H{"response": name})

}

func getGoogleInfo(accessToken string) (Email string, Name string, Error error) {
	uri := "https://www.googleapis.com/oauth2/v2/userinfo"

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return "", "", err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}

	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	var response = struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}{}

	err = json.Unmarshal(buf, &response)
	if err != nil {
		return "", "", err
	}

	return response.Email, response.Name, nil
}

func CreateSession(email string) (SessionToken string, Error error) {

	// create a new session token
	sessionToken := uuid.NewString()

	dbname := os.Getenv("POSTGES_DB")
	dbuser := os.Getenv("POSTGRES_USERNAME")
	dbpassword := os.Getenv("POSTGRES_PASSWORD")
	dbhost := os.Getenv("POSTGRES_HOST")

	connString := fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=require port=5432", dbname, dbuser, dbpassword, dbhost)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return "", err
	}
	defer db.Close()

	_, err = db.Exec(`
	INSERT INTO session (email, session_token) VALUES ($1, $2)
	ON CONFLICT (email) DO UPDATE SET session_token = $2
	`, email, sessionToken)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return sessionToken, nil

}

func GetSession(c *gin.Context) {

	token, err := c.Cookie("session_token")
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid token"})
		return
	}

	fmt.Println(token)

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

	var email string

	err = db.QueryRow(`SELECT email FROM session WHERE session_token = $1`, token).Scan(&email)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(email)

	c.JSON(200, gin.H{"response": email})
}
