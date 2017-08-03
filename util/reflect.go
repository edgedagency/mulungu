package util

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//ReflectDetermineKind returns kind, if subject is array type check type kind
func ReflectDetermineKind(subject interface{}) reflect.Kind {
	kind := ReflectKind(subject)
	switch kind {
	case reflect.Slice:
		kind = ElemKind(subject)
	}
	return kind
}

//ReflectType returns kind of interface using reflection
func ReflectType(subject interface{}) reflect.Type {
	return reflect.TypeOf(subject)
}

//ReflectKind returns kind of interface using reflection
func ReflectKind(subject interface{}) reflect.Kind {
	return ReflectType(subject).Kind()
}

//ReflectIsKindJSONNumber determines if json number
func ReflectIsKindJSONNumber(subject interface{}) bool {

	if reflect.TypeOf(subject) == reflect.TypeOf((*new([]json.Number))) {
		return true
	} else if reflect.TypeOf(subject) == reflect.TypeOf((*new(json.Number))) {
		return true
	} else if reflect.TypeOf(subject) == reflect.TypeOf(([]interface{})(nil)) {
		if reflect.TypeOf(subject.([]interface{})[0]) == reflect.TypeOf((*new(json.Number))) {
			return true
		}
	}

	return false
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

//IsBool returns true if interface passed is a bool
func IsBool(subject interface{}) bool {
	if ReflectKind(subject) == reflect.Bool {
		return true
	}
	return false
}

//ElemKind returns kind of type elem
func ElemKind(subject interface{}) reflect.Kind {
	var elemKind reflect.Kind
	//json.Number in slices panic, this is to help skip
	if reflect.TypeOf(subject) == reflect.TypeOf((*new([]json.Number))) {
		elemKind = reflect.String
	} else if reflect.TypeOf(subject) == reflect.TypeOf([]interface{}(nil)) {
		elemKind = reflect.TypeOf(subject).Elem().Kind()
		//FIXME: make recursive
		if elemKind == reflect.Interface {
			//inconclusive, dig deeper
			elemKind = ReflectKind(subject.([]interface{})[0])
		}

	}

	return elemKind
}

//ElemKindIsString determins is element type is a string
func ElemKindIsString(subject interface{}) bool {
	if ElemKind(subject) == reflect.String {
		return true
	}
	return false
}

//AssertString assert interface to string
func AssertString(subject interface{}) string {
	if IsString(subject) {
		return subject.(string)
	}
	panic(fmt.Errorf("subject %#v not a string", subject))
}

//AssertBool assert interface to string
func AssertBool(subject interface{}) bool {
	if IsBool(subject) {
		return subject.(bool)
	}
	panic(fmt.Errorf("subject %#v not a bool", subject))
}

//AssertStringSlice assert interface to string
func AssertStringSlice(subject interface{}) []string {
	if IsSlice(subject) {
		if ElemKindIsString(subject) {
			return subject.([]string)
		}
	}
	panic(fmt.Errorf("subject %#v not a bool", subject))
}
