package util

import "reflect"

//ToCents returns value as cents
func ToCents(subject interface{}) interface{} {
	switch ReflectKind(subject) {
	case reflect.Int:
		return subject.(int) * 100
	case reflect.String:
		return ToCents(NumberizeString(subject))
	case reflect.Int16:
		return subject.(int16) * 100
	case reflect.Int32:
		return subject.(int32) * 100
	case reflect.Int64:
		return subject.(int64) * 100
	case reflect.Float32:
		return subject.(float32) * 100
	case reflect.Float64:
		return subject.(float64) * 100
	}
	return subject
}

//FromCents returns value from cents to non cents
func FromCents(subject interface{}) interface{} {
	switch ReflectKind(subject) {
	case reflect.Int:
		return subject.(int) / 100
	case reflect.String:
		return ToCents(NumberizeString(subject))
	case reflect.Int16:
		return subject.(int16) / 100
	case reflect.Int32:
		return subject.(int32) / 100
	case reflect.Int64:
		return subject.(int64) / 100
	case reflect.Float32:
		return subject.(float32) / 100
	case reflect.Float64:
		return subject.(float64) / 100
	}
	return subject
}
