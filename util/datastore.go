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

	if kind != reflect.Slice {
		fmt.Println("subject not a slice")
		return IsDatastoreAcceptableKind(subject, false)
	}
	fmt.Println("subject is a slice/map/array type")
	return IsDatastoreAcceptableKind(subject, true)
}

//IsDatastoreAcceptableKind checks if type and type elements are acceptable to datastore
func IsDatastoreAcceptableKind(subject interface{}, multiple bool) bool {
	kind := ReflectDetermineKind(subject)

	fmt.Println("IsDatastoreAcceptableKind : Kind " + kind.String())

	switch kind {
	case reflect.Bool, reflect.String, reflect.Int, reflect.Int8,
		reflect.Int16, reflect.Int32, reflect.Int64, reflect.Slice,
		reflect.Interface, reflect.Array, reflect.Float32, reflect.Float64:
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
