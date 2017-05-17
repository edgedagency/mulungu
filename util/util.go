package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"reflect"

	"golang.org/x/crypto/bcrypt"
)

//InterfaceToByte converts abitary interface to byte
func InterfaceToByte(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

//MapToJSONString convert map to string to string
func MapToJSONString(subject map[string]string) string {

	bytes, err := json.Marshal(subject)
	if err != nil {
		return ""
	}
	return string(bytes)
}

//MapInterfaceToJSONString convert map to string to string
func MapInterfaceToJSONString(subject map[string]interface{}) string {
	bytes, err := json.Marshal(subject)
	if err != nil {
		return ""
	}
	return string(bytes)
}

//JSONStringToMap convert string map[string]string
func JSONStringToMap(subject string) map[string]string {
	outputMap := make(map[string]string)
	err := json.Unmarshal([]byte(subject), &outputMap)
	if err != nil {
		return nil
	}
	return outputMap
}

// JSONDecode converts bytes to map[string]interface{} specified
func JSONDecode(b []byte) (map[string]interface{}, error) {
	results := make(map[string]interface{})
	err := json.Unmarshal(b, &results)

	if err != nil {
		return nil, errors.New("Failed to decode data")
	}

	return results, nil
}

// JSONDecodeHTTPRequest Unmarshal http.Request.Body to interface
func JSONDecodeHTTPRequest(r *http.Request) (map[string]interface{}, error) {
	defer r.Body.Close()
	return ToMapStringInterface(r.Body)
}

// JSONDecodeHTTPResponse Unmarshal http.Request.Body to interface
func JSONDecodeHTTPResponse(r *http.Response) (map[string]interface{}, error) {
	defer r.Body.Close()
	return ToMapStringInterface(r.Body)
}

//ToMapStringInterface converts io.Reader to map[string]interface
func ToMapStringInterface(r io.Reader) (map[string]interface{}, error) {
	results := make(map[string]interface{})
	decoder := json.NewDecoder(r)
	decoder.UseNumber()
	decodeErr := decoder.Decode(&results)

	switch {
	case decodeErr == io.EOF:
		fmt.Println("request has no body, decoding skipped returning nil")
		return nil, nil
	case decodeErr != nil:
		return nil, fmt.Errorf("Failed to decode reader, error %s", decodeErr.Error())
	}

	return results, nil
}

//MD5Hash Hash encode string
func MD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

//ReflectKind returns kind of interface using reflection
func ReflectKind(subject interface{}) reflect.Kind {
	typeKind := reflect.ValueOf(subject).Type().Kind()
	return typeKind
}

//IsString returns true if interface passed is an array
func IsString(subject interface{}) bool {
	if ReflectKind(subject) == reflect.String {
		return true
	}
	return false
}

//IsArray returns true if interface passed is an array
func IsArray(subject interface{}) bool {
	if ReflectKind(subject) == reflect.Array {
		return true
	}
	return false
}

//IsSlice returns true if interface passed is a slice
func IsSlice(subject interface{}) bool {
	if ReflectKind(subject) == reflect.Slice {
		return true
	}
	return false
}

//IsMap returns true if interface passed is a map
func IsMap(subject interface{}) bool {
	if ReflectKind(subject) == reflect.Map {
		return true
	}
	return false
}

//encryption.util

//ComparePlainAndHashed compares a plain non encrypted byte to encrytpted byte
func ComparePlainAndHashed(subject, subjectHashed []byte) (same bool, err error) {
	err = bcrypt.CompareHashAndPassword(subjectHashed, subject)
	if err != nil {
		return false, err
	}
	return true, err
}

//StringTobyte converts string t byte
func StringTobyte(subject string) []byte {
	return []byte(subject)
}

//Search searchs a subject for element
func Search(key, subject interface{}) interface{} {
	var searchResult interface{}

	//log.Debugf("deep searching, key %s", key)
	switch ReflectKind(subject) {
	case reflect.Slice:
		//log.Debug("subject is a slice")
		subjectSlice := subject.([]interface{})
		//log.Debugf("slice size %d", len(subjectSlice))
		for _, sliceValue := range subjectSlice {
			//log.Debugf("deep searching slice @ index %d", index)
			searchResult = Search(key, sliceValue)
		}

		break
	case reflect.Map:
		//log.Debug("subject is a map")
		subjectMap := subject.(map[string]interface{})
		for index, value := range subjectMap {
			//log.Debugf("iterating map keys, @ key %v value %v", index, value)
			//maps can have a varse number of values array, slice, map, so their values are subjected to possibly recursive
			if index == key {
				//log.Debugf("found value %v @ index %v", value, index)
				searchResult = value
			}
		}
		break
	}

	return searchResult
}

//InterfaceToString convets interface{} to string
func InterfaceToString(i interface{}) string {
	return i.(string)
}

//InterfaceToStringSlice converts []interface{} to []string
func InterfaceToStringSlice(i interface{}) []string {
	interfaceSlice := i.([]interface{})
	stringSlice := make([]string, len(interfaceSlice))
	for index, value := range interfaceSlice {
		stringSlice[index] = InterfaceToString(value)
	}
	return stringSlice
}

//InterfaceToMapString converts interface to map[string]string
func InterfaceToMapString(subject interface{}) map[string]string {
	switch ReflectKind(subject) {
	case reflect.String:
		return JSONStringToMap(subject.(string))
	}

	return nil
}

//GenerateRandomCode generates random codes based on provided characters of internal character set
func GenerateRandomCode(length int, characters string) string {
	if characters == "" {
		characters = "ohruix3yetu5dei7oqu4gothah4Esei6xudez9saejueshuThaj4ooPh1Shi8engahGhiesaeng9meib8iPhaeNg7eikohSh8ae9"
	}

	b := make([]byte, length)
	for i := range b {
		b[i] = characters[rand.Int63()%int64(len(characters))]
	}
	return string(b)
}
