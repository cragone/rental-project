package handlers

import (
	"database/sql"
	"fmt"
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

	googleURI := fmt.Sprintf(`https://accounts.google.com/o/oauth2/auth?client_id=%s&redirect_uri=%s&response_type=code&scope=openid+profile+email&prompt=consent`, os.Getenv("CLIENT_ID"), os.Getenv("REDIRECT_URI"))

	c.JSON(200, gin.H{"response": googleURI})
}
