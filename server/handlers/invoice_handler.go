package handlers

import (
	"server/invoice"

	"github.com/gin-gonic/gin"
)

func HandleNewOrder(c *gin.Context) {

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
