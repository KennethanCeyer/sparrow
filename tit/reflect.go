package tit

import (
	"reflect"
	"tit/utils"
)

func getFieldNames(t reflect.Type) []string {
	numField := t.NumField()
	names := make([]string, numField)
	len := 0
	for i:=0; i<numField; i++ {
		name := t.Field(i).Name
		if name[0] < 'A' || name[0] > 'Z' {
			continue
		}

		names[i] = name
		len++
	}
	return names[:len]
}

func getFieldTags(t reflect.Type) []string {
	numField := t.NumField()
	tag := make([]string, numField)
	for i:=0; i<numField; i++ {
		tag[i] = t.Field(i).Tag.Get(AppName)
	}
	return tag
}

func getFieldByTag(v reflect.Value, tag string) reflect.Value {
	t := v.Type()
	numField := t.NumField()
	for i:=0; i<numField; i++ {
		if tag == t.Field(i).Tag.Get(AppName) {
			return v.Field(i)
		}
	}

	var emptyValue reflect.Value
	return emptyValue
}

func findPropertyByKey(key string, properties []string) string {
	for _, value := range properties {
		filterKeys := []string{key, utils.ToCamelCase(key), utils.ToPascalCase(key)}
		for _, filterValue := range filterKeys {
			if filterValue == value {
				return filterValue
			}
		}
	}
	return ""
}

func propagateToModel(data map[interface{}]interface{}, model interface{}) {
	v := reflect.ValueOf(model)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			v.Set(reflect.New(reflect.TypeOf(v.Elem())))
		}
		v = v.Elem()
	}

	typeSt := v.Type()
	names := getFieldNames(typeSt)
	tags := getFieldTags(typeSt)

	for key, value := range data {
		if utils.InArray(key, tags) {
			field := getFieldByTag(v, key.(string))
			field.Set(reflect.ValueOf(value))
			continue
		}

		matchKey := findPropertyByKey(key.(string), names)
		if matchKey == "" {
			continue
		}

		v.FieldByName(matchKey).Set(reflect.ValueOf(value))
	}
}
