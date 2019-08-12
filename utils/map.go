package utils

import "reflect"

func Flatten(data map[interface{}]interface{}) map[interface{}]interface{} {
	flatten := make(map[interface{}]interface{})
	for key, value := range data {
		if reflect.ValueOf(value).Kind() == reflect.Map {
			interfaceMap := value
			if reflect.TypeOf(value) == reflect.TypeOf(map[string]interface{}{}) {
				interfaceMap = ConvertMapForm(value.(map[string]interface{}))
			}

			subData := Flatten(interfaceMap.(map[interface{}]interface{}))
			for subKey, subValue := range subData {
				if _, ok := flatten[subKey]; !ok {
					flatten[subKey] = subValue
				}
			}

			continue
		}

		if _, ok := flatten[key]; !ok {
			flatten[key] = value
		}
	}
	return flatten
}

func ConvertMapForm(data map[string]interface{}) map[interface{}]interface{} {
	dataPack := make(map[interface{}]interface{})
	for key, value := range data {
		if reflect.ValueOf(value).Kind() == reflect.Map {
			value = ConvertMapForm(value.(map[string]interface{}))
		}
		dataPack[key] = value
	}
	return dataPack
}
