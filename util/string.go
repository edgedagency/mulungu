package util

import (
	"reflect"
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

//ToString Converts interface to a string
func ToString(subject interface{}) string {

	switch ReflectKind(subject) {
	case reflect.String:
		return subject.(string)
	case reflect.Int64:
		return strconv.FormatInt(subject.(int64), 10)
	case reflect.Float64:
		return strconv.FormatFloat(subject.(float64), 'E', -1, 64)
	}

	return ""
}
