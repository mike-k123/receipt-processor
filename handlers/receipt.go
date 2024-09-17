package handlers

import (
	"net/http"
	"receipt-processor/models"
	"receipt-processor/services"
	"receipt-processor/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var receiptStore = storage.NewInMemoryStore()

// Handler to process a receipt
func ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	receiptID := uuid.New().String()

	points := services.CalculatePoints(receipt)

	receiptStore.Save(receiptID, points)

	c.JSON(http.StatusOK, gin.H{"id": receiptID})
}

// Handler to get points for a receipt
func GetPoints(c *gin.Context) {
	id := c.Param("id")

	points, found := receiptStore.Get(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}
