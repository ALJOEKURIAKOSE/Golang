package main

import (
	"net/http"
	"strconv"
	"unicode"

	"github.com/gin-gonic/gin"
)

func ValidateCreditCard(number string) bool {
	// Remove all spaces from the number
	number = removeSpaces(number)
	// Check if the number contains only digits
	if !isNumeric(number) {
		return false
	}
	// Apply Luhn algorithm
	return luhnCheck(number)
}

// removeSpaces removes all spaces from the input string.
func removeSpaces(input string) string {
	output := ""
	for _, char := range input {
		if !unicode.IsSpace(char) {
			output += string(char)
		}
	}
	return output
}

// isNumeric checks if a string contains only digits.
func isNumeric(input string) bool {
	for _, char := range input {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

// luhnCheck validates a credit card number using the Luhn algorithm.
func luhnCheck(number string) bool {
	sum := 0
	double := false

	// Process each digit starting from the rightmost side
	for i := len(number) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(number[i]))
		if err != nil {
			return false
		}

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		double = !double
	}

	return sum%10 == 0
}

func main() {
	r := gin.Default()

	// Endpoint to validate credit card number
	r.POST("/validate", func(c *gin.Context) {
		var request struct {
			CardNumber string `json:"card_number"`
		}

		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		if ValidateCreditCard(request.CardNumber) {
			c.JSON(http.StatusOK, gin.H{"valid": true})
		} else {
			c.JSON(http.StatusOK, gin.H{"valid": false})
		}
	})

	// Start the server on port 8080
	r.Run(":8080")
}
