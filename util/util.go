package util

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
)

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
	results := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&results)
	if err != nil {
		return nil, errors.New("Failed to decode http request")
	}
	return results, nil
}

// JSONDecodeHTTPResponse Unmarshal http.Request.Body to interface
func JSONDecodeHTTPResponse(r *http.Response) (map[string]interface{}, error) {
	defer r.Body.Close()
	results := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&results)
	if err != nil {
		return nil, errors.New("Failed to decode http response")
	}
	return results, nil
}

//MD5Hash Hash encode string
func MD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

//BasicAuthUsernamePassword returns username and password from basic auth base64 encoded string
func BasicAuthUsernamePassword(base64Encoded string) (username, password string, err error) {
	if len(base64Encoded) <= 0 {
		return "", "", errors.New("empty string provided")
	}
	data, err := base64.StdEncoding.DecodeString(base64Encoded)
	if err != nil {
		log.Fatal("error:", err)
		return "", "", errors.New("Failed to decode string")
	}
	usernameAndPassword := strings.Split(string(data), ":")

	return usernameAndPassword[0], usernameAndPassword[1], nil
}
