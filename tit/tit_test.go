package tit

import (
	"fmt"
	"github.com/KennethanCeyer/tit/resolver"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetResolver(t *testing.T) {
	var expected *resolver.Resolver
	handler, err := getResolver(resolver.YAMLType)

	assert.NoError(t, err)
	assert.Equal(t, reflect.TypeOf(expected).String(), reflect.TypeOf(handler).String())
}

func TestGetResolverWithUnknownType(t *testing.T) {
	typeName := "unknown"
	handler, err := getResolver(resolver.Type(typeName))

	assert.Error(t, err)
	assert.EqualError(t, err, fmt.Sprintf("`%v` is not supported type of resolver", typeName))
	assert.Nil(t, handler)
}

func TestReadFileWithYAML(t *testing.T) {
	type Config struct {
		AppName    string
		AppVersion string
		Debug      bool
	}
	var config Config

	err := ReadFile("../testdata/config.yml", &config)

	assert.NoError(t, err)
	assert.Equal(t, "tit", config.AppName)
	assert.Equal(t, "1.0.0", config.AppVersion)
	assert.Equal(t, true, config.Debug)
}

func TestReadFileWithJSON(t *testing.T) {
	type Config struct {
		AppName    string
		AppVersion string
		Debug      bool
	}
	var config Config

	err := ReadFile("../testdata/config.json", &config)

	assert.NoError(t, err)
	assert.Equal(t, "tit", config.AppName)
	assert.Equal(t, "1.0.0", config.AppVersion)
	assert.Equal(t, true, config.Debug)
}

func TestReadFileWithTOML(t *testing.T) {
	type Config struct {
		AppName    string
		AppVersion string
		Debug      bool
	}
	var config Config

	err := ReadFile("../testdata/config.toml", &config)

	assert.NoError(t, err)
	assert.Equal(t, "tit", config.AppName)
	assert.Equal(t, "1.0.0", config.AppVersion)
	assert.Equal(t, true, config.Debug)
}

func TestReadFileWithMapModel(t *testing.T) {
	t.Skip()
	var config map[string]interface{}

	err := ReadFile("../testdata/config.yml", &config)

	assert.NoError(t, err)
	assert.Equal(t, "tit", config["appName"])
	assert.Equal(t, "1.0.0", config["appVersion"])
	assert.Equal(t, true, config["debug"])
}
