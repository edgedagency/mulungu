package mulungu

import (
	"encoding/json"
	"net/http"
)

//RequestData represents request data
type RequestData map[string]interface{}

//NewRequestData returns new request data object, this can be used to manipulate request data. Form fields e.t.c.
func NewRequestData(r *http.Request) *RequestData {
	rd := &RequestData{}
	rd.Hydrate(r)

	return rd
}

//Hydrate hydrates request based on data received from http.request
func (rd *RequestData) Hydrate(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(rd)
	if err != nil {
		return err
	}
	return nil
}

//Get returns item by key
func (rd *RequestData) Get(key string) interface{} {
	if item, ok := (*rd)[key]; ok {
		return item
	}
	return nil
}

//Set sets a new value for map
func (rd *RequestData) Set(key string, value interface{}) {
	(*rd)[key] = value
}
