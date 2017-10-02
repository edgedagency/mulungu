package util

import (
	"encoding/base64"
	"reflect"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
)

//StringDecode decodes base64 encoded string content
func StringDecode(encodedString string) ([]byte, error) {
	decoded, decodedError := base64.StdEncoding.DecodeString(encodedString)
	if decodedError != nil {
		return nil, decodedError
	}
	return decoded, nil
}

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
	case reflect.Int:
		return strconv.Itoa(subject.(int))
	case reflect.Int64:
		return strconv.FormatInt(subject.(int64), 10)
	case reflect.Float64:
		return strconv.FormatFloat(subject.(float64), 'E', -1, 64)
	}

	return ""
}

func NumberizeString(subject interface{}) interface{} {
	if govalidator.IsInt(subject.(string)) {
		return StringToInt(subject.(string))
	} else if govalidator.IsFloat(subject.(string)) {
		return StringToInt64(subject.(string))
	}
	return subject
}

//StringTobyte converts string t byte
func StringTobyte(subject string) []byte {
	return []byte(subject)
}
