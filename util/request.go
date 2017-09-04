package util

import (
	"net/http"

	"github.com/gorilla/context"
)

// type RequestHandler struct {
// 	Request *http.Request
// }
//
// func (r *RequestHandler) Context() context.Context {
//
// }

//ContextSetValue set a value in context
func ContextSetValue(r *http.Request, key interface{}, val interface{}) {
	context.Set(r, key, val)
}

//ContextGetValue get a value from context
func ContextGetValue(r *http.Request, key interface{}) interface{} {
	return context.Get(r, key)
}

//ContextGetAll return all values in context
func ContextGetAll(r *http.Request) map[interface{}]interface{} {
	return context.GetAll(r)
}
