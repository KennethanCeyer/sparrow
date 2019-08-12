package handler

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestJSONResolver_Resolve(t *testing.T) {
	filePath := "../../testdata/config.json"
	data, err := ioutil.ReadFile(filePath)
	assert.NoError(t, err)

	handler := new(JSONResolver)
	dataPack, err := handler.Resolve(data)

	assert.NoError(t, err)
	assert.Equal(t, 3, len(dataPack))
	assert.Equal(t, "tit", dataPack["appName"])
	assert.Equal(t, "1.0.0", dataPack["appVersion"])
	assert.Equal(t, true, dataPack["debug"])
}
