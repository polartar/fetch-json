package models
import (
	 
	"github.com/google/uuid"
	"errors"
	"fmt"
)

type Item struct {
	ShortDescription        string      `json:"shortDescription"`
	Price    				string      `json:"price"`
}

 
type ReceiptInput struct {
	Retailer        string      `json:"retailer"`
	PurchaseDate    string      `json:"purchaseDate"`
	PurchaseTime    string      `json:"purchaseTime"`
	Items      		[]Item  	`json:"items"`  
	Total    	    string      `json:"total"`
}

type Receipt struct {
	ReceiptInput
	ID          	string      `json:"id"`
}

// Custom Error Response Struct
type ErrorResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

type PointsResponse struct {
    Points int `json:"points"`
}

// Failure response with a message
type PointsErrorResponse struct {
    Message string `json:"message"`
}

var receiptList []Receipt

 
// Saves a receipt to the list
func Save(input ReceiptInput) (*Receipt) {
	receiptID := uuid.New().String()
	fmt.Println("Save method called")
    fmt.Println("Generated receipt ID:", receiptID)
	receipt := Receipt{
		ReceiptInput: input, 
		ID:           receiptID,
	}

	receiptList = append(receiptList, receipt)

	return &receipt
}
 
// Fetches a receipt from the list
func FetchReceipt(id string) (Receipt, error) {
	
	for _, receipt := range receiptList {
		if receipt.ID == id {
			return receipt, nil
		}
	}

	return Receipt{}, errors.New("receipt not found")
}
 