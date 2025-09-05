package validators

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

const (
	emptyString = ""
	emptyInt    = 0
)

func ValidateMonth(c *gin.Context) (bool, int) {
	monthStr := c.Query("month")
	if monthStr == "" {
		monthStr = "0"
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		return false, 0
	}

	if month < 1 {
		month = int(time.Now().Month())
	}

	if month < 1 || month > 12 {
		return false, 0
	}

	return true, month
}

func ValidateYear(c *gin.Context) (bool, int) {
	currentYear := time.Now().Year()

	yearStr := c.Query("year")
	if yearStr == "" {
		yearStr = "0"
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return false, 0
	}

	if year < 1 {
		year = currentYear
	}

	if year > currentYear {
		return false, 0
	}

	return true, year
}
