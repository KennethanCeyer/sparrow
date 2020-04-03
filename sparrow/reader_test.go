package tit

import (
	"github.com/KennethanCeyer/tit/resolver"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleansingAliasWithYML(t *testing.T) {
	fileName := "yml"
	fileType := cleansingAlias(fileName)

	assert.Equal(t, resolver.YAMLType, fileType)
}

func TestCleansingAliasWithYAML(t *testing.T) {
	fileName := "yaml"
	fileType := cleansingAlias(fileName)

	assert.Equal(t, resolver.YAMLType, fileType)
}

func TestCleansingAliasWithJSON(t *testing.T) {
	fileName := "json"
	fileType := cleansingAlias(fileName)

	assert.Equal(t, resolver.JSONType, fileType)
}

func TestCleansingAliasWithTOML(t *testing.T) {
	fileName := "toml"
	fileType := cleansingAlias(fileName)

	assert.Equal(t, resolver.TOMLType, fileType)
}

func TestGetTypeFromExtWithYML(t *testing.T) {
	fileName := "config.yml"
	fileType, err := getTypeFromExt(fileName)

	assert.NoError(t, err)
	assert.Equal(t, resolver.YAMLType, fileType)
}

func TestGetTypeFromExtWithYAML(t *testing.T) {
	fileName := "config.yaml"
	fileType, err := getTypeFromExt(fileName)

	assert.NoError(t, err)
	assert.Equal(t, resolver.YAMLType, fileType)
}

func TestGetTypeFromExtWithJSON(t *testing.T) {
	fileName := "config.json"
	fileType, err := getTypeFromExt(fileName)

	assert.NoError(t, err)
	assert.Equal(t, resolver.JSONType, fileType)
}

func TestGetTypeFromExtWithTOML(t *testing.T) {
	fileName := "config.toml"
	fileType, err := getTypeFromExt(fileName)

	assert.NoError(t, err)
	assert.Equal(t, resolver.TOMLType, fileType)
}
