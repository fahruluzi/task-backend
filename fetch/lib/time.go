package lib

import (
	"strings"
	"time"
)

// ParseDate parse date string to format 2006-01-02
func ParseDate(date string) time.Time {
	if strings.Contains(date, " ") {
		split := strings.Split(date, " ")
		date = split[0]
	} else if strings.Contains(date, "T") {
		split := strings.Split(date, "T")
		date = split[0]
	}

	parsedDate, _ := time.Parse("2006-01-02", date)
	return parsedDate
}
