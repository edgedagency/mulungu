package util

import (
	"encoding/json"
	"net/http"
	"strings"

	"google.golang.org/appengine/log"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

//WriteJSON outputs json to response writer and sets up the right mimetype
func WriteJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

//WriteXML outputs interface to xml
func WriteXML(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	w.WriteHeader(statusCode)
	b, err := MapToXML(data.(map[string]interface{}))
	if err != nil {
		w.Write([]byte("<error>failed to process data</error>"))
	}

	w.Write(b)
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

//SetEnvironmentOnNamespace attaches environment in request to provided spacename
func SetEnvironmentOnNamespace(ctx context.Context, namespace string, r *http.Request) string {
	environment := r.URL.Query().Get("env")

	if strings.HasPrefix(namespace, environment) == false {
		if environment != "" {
			log.Debugf(ctx, "setting environment on environment:%s", environment)
			namespace = strings.Join([]string{environment, namespace}, ".")
		} else {
			log.Debugf(ctx, "no environment specified for request, e.g. use http://example.com?env=dev to set environment to dev")
		}
	} else {
		log.Debugf(ctx, "skipped setting environment namespace, namespace prefixed with:%s namespace:%s", environment, namespace)
	}
	return namespace
}

//SetNamespace  attaches environment in request to provided spacename
func SetNamespace(ctx context.Context, r *http.Request) string {
	namespace := r.URL.Query().Get("ns")

	if namespace != "" {
		return namespace
	}

	return ""
}
