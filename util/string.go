package util

import (
	"strconv"
	"strings"
)

//StringContains checks if a string contains any of the pradicates send
func StringContains(subject string, pradicates []string) bool {
	for _, pradicate := range pradicates {
		if strings.Contains(subject, pradicate) {
			return true
		}
	}
	return false
}

//StringToInt converts a string to an integer
func StringToInt(subject string) int {
	i, err := strconv.Atoi(subject)
	if err != nil {
		return 0
	}
	return i
}

//StringToInt64 converts a string to an integer
func StringToInt64(subject string) int64 {
	i, err := strconv.ParseInt(subject, 10, 64)
	if err != nil {
		return 0
	}
	return i
}
