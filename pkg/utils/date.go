package utils

import "time"

func GetCurrentDay() int {
	return getCurrentTime().Day()
}

func GetCurrentYear() int {
	return getCurrentTime().Year()
}

func getCurrentTime() time.Time {
	return time.Now()
}
