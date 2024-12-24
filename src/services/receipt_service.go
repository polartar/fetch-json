package services


import (

	"github.com/polartar/fetch-test/src/models"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)


func CalculatePoints(receipt models.Receipt) int {
	points := 0

	points += countAlphanumeric(receipt.Retailer)

	total, err := strconv.ParseFloat(receipt.Total, 64)
	fmt.Println("total:", points)
	if err != nil {
		fmt.Println("Error parsing total:", err)
		return 0
	}
	if total == float64(int(total)) {
		points += 50
	}

	if math.Mod(total, 0.25) == 0 {
		points += 25
	}
	fmt.Println("total:", points)

	points += 5 * (len(receipt.Items) / 2)
	fmt.Println("total:", points)
	for _, item := range receipt.Items {
		itemPrice, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			fmt.Println("Error parsing item price:", err)
			continue
		}

		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			points += int(math.Ceil(itemPrice * 0.2))
		}
	}
	fmt.Println("total:", points)
	if total > 10.00 {
		points += 5
	}
	fmt.Println("total:", points)
	parsedDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err == nil && parsedDate.Day()%2 != 0 {
		points += 6
	}
	fmt.Println("total:", points)
	parsedTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err == nil && parsedTime.Hour() >= 14 && parsedTime.Hour() < 16 {
		points += 10
	}
	fmt.Println("total:", points)
	return points
}

func countAlphanumeric(input string) int {
	count := 0
	for _, char := range input {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			count++
		}
	}
	return count
}