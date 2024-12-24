package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/polartar/fetch-test/src/models"
	"github.com/polartar/fetch-test/src/services"
	 
  
)

// @Summary Get points11 for a receipt by ID
// @Description Fetch the receipt by its ID and calculate the points based on predefined rules
// @Tags receipts
// @Accept json
// @Produce json
// @Param id path string true "Receipt ID"
// @Success 200 {object} models.PointsResponse "Points calculated successfully"
// @Failure 400 {object} models.PointsErrorResponse "Bad request, receipt ID is required or receipt not found"
// @Router /receipts/{id}/points [get]
func GetReceiptByID(c *gin.Context) {
	receiptId := c.Param("id")
	if receiptId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Receipt ID is required"})
		return
	}
	receipt, err := models.FetchReceipt(receiptId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),  
		})
		return
	}

	points := services.CalculatePoints(receipt)


	c.JSON(http.StatusOK,   gin.H{
		"points":  points,
	})
}


// @Summary Create a new receipt
// @Description Create a new receipt by providing the retailer, purchase date, time, items, and total amount
// @Tags receipts
// @Accept json
// @Produce json
// @Param receiptInput body models.ReceiptInput true "Receipt input data"
// @Success 201 {object} models.Receipt "Receipt created successfully"
// @Failure 400 {object} models.ErrorResponse "Bad request, invalid input"
// @Router /receipts/process [post]
func CreateReceipt(context *gin.Context) {

	var input models.ReceiptInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Invalid input: " + err.Error(),
			"data":    nil,
		})
		return
	}

	receipt := models.Save(input)

	context.JSON(http.StatusCreated, receipt)
}
 