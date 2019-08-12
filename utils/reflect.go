package utils

import (
	"reflect"
)


func SoftSetFloat(field reflect.Value, value reflect.Value) {
	switch value.Interface().(type) {
	case int, int8, int16, int32, int64:
		field.SetFloat(float64(value.Int()))
	case uint, uint8, uint16, uint32, uint64:
		field.SetFloat(float64(value.Uint()))
	default:
		field.SetFloat(value.Float())
	}
}

func SoftSetInt(field reflect.Value, value reflect.Value) {
	switch value.Interface().(type) {
	case float32, float64:
		field.SetInt(int64(value.Float()))
	case uint, uint8, uint16, uint32, uint64:
		field.SetInt(int64(value.Uint()))
	default:
		field.SetInt(value.Int())
	}
}

func SoftSetUint(field reflect.Value, value reflect.Value) {
	switch value.Interface().(type) {
	case float32, float64:
		field.SetUint(uint64(value.Float()))
	case int, int8, int16, int32, int64:
		field.SetUint(uint64(value.Int()))
	default:
		field.SetUint(value.Uint())
	}
}

func SoftSet(field reflect.Value, value reflect.Value) {
	switch field.Kind() {
	case reflect.Float32, reflect.Float64:
		SoftSetFloat(field, value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		SoftSetInt(field, value)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		SoftSetUint(field, value)
	default:
		field.Set(value)
	}
}
