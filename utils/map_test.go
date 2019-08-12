package utils

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestConvertMapForm(t *testing.T) {
	data := map[string]interface{}{
		"apple": 1,
		"banana": "yellow",
		"grape": true,
	}
	convertData := ConvertMapForm(data)

	assert.Equal(t, "int", reflect.TypeOf(convertData["apple"]).String())
	assert.Equal(t, "string", reflect.TypeOf(convertData["banana"]).String())
	assert.Equal(t, "bool", reflect.TypeOf(convertData["grape"]).String())

	for key := range data {
		assert.Equal(t, "string", reflect.TypeOf(key).String())
	}
}
