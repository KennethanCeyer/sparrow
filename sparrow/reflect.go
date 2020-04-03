package sparrow

import (
	"github.com/KennethanCeyer/sparrow/utils"
	"reflect"
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

func processStructValue(model reflect.Value, key string, value interface{}, names []string, tags []string) {
	var field reflect.Value

	if utils.InArray(key, tags) {
		field = getFieldByTag(model, key)
	} else if matchKey := findPropertyByKey(key, names); matchKey != "" {
		field = model.FieldByName(matchKey)
	}

	if reflect.ValueOf(value).Kind() == reflect.Map {
		propagateToModel(value.(map[interface{}]interface{}), field)
		return
	}

	if !field.CanSet() {
		return
	}

	utils.SoftSet(field, reflect.ValueOf(value))
}

func propagateToStruct(data map[interface{}]interface{}, model reflect.Value) {
	typeSt := model.Type()

	names := getFieldNames(typeSt)
	tags := getFieldTags(typeSt)

	for resolvedKey, resolvedValue := range data {
		processStructValue(model, resolvedKey.(string), resolvedValue, names, tags)
	}
}

func propagateToMap(data map[interface{}]interface{}, model reflect.Value) {
	model.Set(reflect.MakeMap(model.Type()))
	for resolvedKey, resolvedValue := range data {
		model.SetMapIndex(reflect.ValueOf(resolvedKey), reflect.ValueOf(resolvedValue))
	}
}


func propagateToModel(data map[interface{}]interface{}, model interface{}) {
	v := reflect.ValueOf(model)

	if v.Type() == reflect.TypeOf(reflect.Value{}) {
		v = model.(reflect.Value)
	}

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			v.Set(reflect.New(reflect.TypeOf(v.Elem())))
		}
		v = v.Elem()
	}

	switch v.Kind(){
	case reflect.Struct:
		propagateToStruct(data, v)
	case reflect.Map:
		propagateToMap(data, v)
	}
}
