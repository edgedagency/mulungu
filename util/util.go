package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// JSONDecode converts bytes to map[string]interface{} specified
func JSONDecode(b []byte) interface{} {
	results := make(map[string]interface{})
	err := json.Unmarshal(b, &results)
	if err != nil {
		log.Println("failed to unmarshal object")
		return nil
	}
	fmt.Println(results)
	return results
}

// JSONDecodeHTTPResponse Unmarshal http.Reponse.Body to interface
func JSONDecodeHTTPResponse(res *http.Response) interface{} {
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("failed to read response body")
	}
	return JSONDecode(b)
}
