package utils

import "time"

func validateDay(day int) bool {
	return day >= 1 && day <= 25
}

func validateYear(year int, currentYear int) bool {
	return year >= 2015 && year <= currentYear
}

func ValidateDayYearInput(day int, year int) bool {
	now := time.Now()

	if year == now.Year() {
		return validateDay(day) && day <= now.Day()
	}

	return validateDay(day) && validateYear(year, now.Year())
}
