package services

import (
	"math"
	"receipt-processor/models"
	"receipt-processor/utils"
	"strconv"
	"strings"
	"time"
)

// CalculatePoints calculates the points for a given receipt
func CalculatePoints(receipt models.Receipt) int {
	points := 0

	retailerName := strings.ReplaceAll(receipt.Retailer, " ", "")
	for _, char := range retailerName {
		if utils.IsAlphanumeric(char) {
			points++
		}
	}

	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == float64(int(total)) {
		points += 50
	}

	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5

	for _, item := range receipt.Items {
		trimmedDesc := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDesc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if date.Day()%2 != 0 {
		points += 6
	}

	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}
