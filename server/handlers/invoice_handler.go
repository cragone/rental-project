package handlers

import (
	"fmt"
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
	}

	var tennant = invoice.Tennant{
		TennantID: p.TennantID,
		Rate:      p.Rate,
	}

	tennants := append([]invoice.Tennant{}, tennant)

	invoice.GenerateInvoices(tennants)

	fmt.Println(tennant)

}

func HandleNewOrderTest(c *gin.Context) {

	orderID, err := invoice.GeneratePaypalOrder(320)
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
