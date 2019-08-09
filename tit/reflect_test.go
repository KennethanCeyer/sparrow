package tit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestModel struct {
	Apple  string `tit:"circle"`
	Banana string
	Grape  string
}

type TestModelWithTag struct {
	Apple  string `tit:"circle"`
	Banana string `tit:long`
	Grape  string `tit:clumpy`
}

func TestPropagateToModelWithoutTag(t *testing.T) {
	testModel := TestModel{}
	data := map[interface{}]interface{} {
		"apple": "red",
		"banana": "yellow",
		"grape": "purple",
	}
	propagateToModel(data, &testModel)

	assert.Equal(t, "red", testModel.Apple)
	assert.Equal(t, "yellow", testModel.Banana)
	assert.Equal(t, "purple", testModel.Grape)
}

func TestPropagateToModelWithTag(t *testing.T) {
	testModel := TestModelWithTag{}
	data := map[interface{}]interface{} {
		"circle": "red",
		"long": "yellow",
		"clumpy": "purple",
	}
	propagateToModel(data, &testModel)

	assert.Equal(t, "red", testModel.Apple)
	assert.Equal(t, "yellow", testModel.Banana)
	assert.Equal(t, "purple", testModel.Grape)
}

func TestPropagateToModelWithNotMatchedTag(t *testing.T) {
	testModel := TestModelWithTag{}
	data := map[interface{}]interface{} {
		"apple": "red",
		"banana": "yellow",
		"grape": "purple",
	}
	propagateToModel(data, &testModel)

	assert.Equal(t, "red", testModel.Apple)
	assert.Equal(t, "yellow", testModel.Banana)
	assert.Equal(t, "purple", testModel.Grape)
}
