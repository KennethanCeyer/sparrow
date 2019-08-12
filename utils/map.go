package utils

func ConvertMapForm(data map[string]interface{}) map[interface{}]interface{} {
	dataPack := make(map[interface{}]interface{})
	for key, value := range data {
		dataPack[key] = value
	}
	return dataPack
}
