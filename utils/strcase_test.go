package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToPascalCase(t *testing.T) {
	fromWhiteSpace := ToPascalCase("from white space")
	fromCamelCase := ToPascalCase("fromCamelCase")
	fromSnakeCase := ToPascalCase("from_snake_case")
	fromKebabCase := ToPascalCase("from-kebab-case")
	fromPascalCase := ToPascalCase("FromPascalCase")

	assert.Equal(t, "FromWhiteSpace", fromWhiteSpace)
	assert.Equal(t, "FromCamelCase", fromCamelCase)
	assert.Equal(t, "FromSnakeCase", fromSnakeCase)
	assert.Equal(t, "FromKebabCase", fromKebabCase)
	assert.Equal(t, "FromPascalCase", fromPascalCase)
}

func TestToCamelCase(t *testing.T) {
	fromWhiteSpace := ToCamelCase("from white space")
	fromCamelCase := ToCamelCase("fromCamelCase")
	fromSnakeCase := ToCamelCase("from_snake_case")
	fromKebabCase := ToCamelCase("from-kebab-case")
	fromPascalCase := ToCamelCase("FromPascalCase")

	assert.Equal(t, "fromWhiteSpace", fromWhiteSpace)
	assert.Equal(t, "fromCamelCase", fromCamelCase)
	assert.Equal(t, "fromSnakeCase", fromSnakeCase)
	assert.Equal(t, "fromKebabCase", fromKebabCase)
	assert.Equal(t, "fromPascalCase", fromPascalCase)
}
