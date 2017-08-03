package util

import (
	"encoding/json"
	"fmt"
	"reflect"

	"cloud.google.com/go/datastore"
)

//IsDatastoreAcceptableType checks if type and type elements are acceptable to datastore
func IsDatastoreAcceptableType(subject interface{}) bool {
	kind := ReflectKind(subject)
	fmt.Println(fmt.Printf("checking is subject is datastore acceptable : subject > %#v kind > %s ", subject, kind.String()))
	return IsDatastoreAcceptableKind(subject)
}

//IsDatastoreAcceptableKind checks if type and type elements are acceptable to datastore
func IsDatastoreAcceptableKind(subject interface{}) bool {
	kind := ReflectDetermineKind(subject)
	fmt.Println("IsDatastoreAcceptableKind : Kind " + kind.String())

	switch kind {
	case reflect.Bool, reflect.String, reflect.Int, reflect.Int8,
		reflect.Int16, reflect.Int32, reflect.Int64, reflect.Slice,
		reflect.Interface, reflect.Array, reflect.Float32, reflect.Float64, reflect.Map:
		return true
	}
	return false
}

//DatastoreConvertJSONNumberToSupportedSlice converts []jsonNumber to []interface
func DatastoreConvertJSONNumberToSupportedSlice(subject interface{}) []interface{} {
	interfaceSlice := make([]interface{}, len(subject.([]interface{})))
	for index, value := range subject.([]interface{}) {
		interfaceSlice[index] = (value.(json.Number)).String()
	}

	return interfaceSlice
}

//GetDatastoreProperty creates a datastore property
func GetDatastoreProperty(name string, index bool, value interface{}) datastore.Property {
	return datastore.Property{Name: name, NoIndex: index, Value: value}
}

//GetDatastoreProperties returns multiple properties
func GetDatastoreProperties(subject []interface{}) []datastore.Property {
	properties := []datastore.Property{}
	for _, item := range subject {
		if ReflectKind(item) == reflect.Map {
			for key, value := range item.(map[string]interface{}) {
				// fmt.Println(fmt.Sprintf("creating value ->>> %s %#v", key, value))
				properties = append(properties, GetDatastoreProperty(key, false, value))
			}
		}
	}
	return properties
}
