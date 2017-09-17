package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/appengine/log"

	"github.com/edgedagency/mulungu/constant"
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

//HTTPGetPath obtains path from request
func HTTPGetPath(r *http.Request) string {
	path := r.URL.Path
	if r.URL.RawQuery != "" {
		path = strings.Join([]string{path, r.URL.RawQuery}, "?")
	}
	return path
}

//HTTPCopyRequestHeader from original request
func HTTPCopyRequestHeader(originalRequest *http.Request, candicateRequest *http.Request) *http.Request {
	for key, headerValue := range originalRequest.Header {
		for _, value := range headerValue {
			candicateRequest.Header.Set(key, value)
		}
	}

	return candicateRequest
}

//HTTPGetGoogleAppEngineServiceURL this returns an app spot host
func HTTPGetGoogleAppEngineServiceURL(r *http.Request, service, defaultServiceHost, path string) string {
	return fmt.Sprintf("%s://%s-dot-%s/%s", "https", service, HTTPGetServiceHost(r, defaultServiceHost), path)
}

//HTTPGetPathVariables returns map[string]string of path varaibles
func HTTPGetPathVariables(r *http.Request) map[string]string {
	return mux.Vars(r)
}

//HTTPGetServiceHost obtains servce host from request headers
func HTTPGetServiceHost(r *http.Request, defaultServiceHost string) string {
	serviceHost := r.Header.Get(constant.HeaderServiceHost)
	if serviceHost == "" {
		return defaultServiceHost
	}
	return serviceHost
}
