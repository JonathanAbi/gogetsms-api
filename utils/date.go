package utils

import (
	"time"
)

func IsLessThan3Days(endDateStr string) bool {
	if len(endDateStr) < 19 {
		return false
	}
	datePart := endDateStr[:19]

	endDate, err := time.Parse("2006-01-02T15:04:05", datePart)
	if err != nil {
		return false
	}

	currentTime := time.Now()

	diff := endDate.Sub(currentTime)
	diffDays := diff.Hours() / 24

	return diffDays >= 0 && diffDays < 3
}
