package util

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//WriteJSON outputs json to response writer and sets up the right mimetype
func WriteJSON(w http.ResponseWriter, responseBody interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(responseBody)
}

//GeneratePath use this if you would like to generate a path based on request, excluding service. Very specific to applications developed for ibudo
func GeneratePath(r *http.Request) string {
	capturedPath := mux.Vars(r)["path"]
	if capturedPath != "" {
		if r.URL.RawQuery != "" {
			return strings.Join([]string{capturedPath, r.URL.RawQuery}, "?")
		}
		return capturedPath
	}
	return ""
}
