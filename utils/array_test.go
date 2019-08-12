package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInArrayWithNumer(t *testing.T) {
	arr := []int{1, 2, 5, 6, 8, 10, 12, 13, 15, 19, 20}
	given := []int{1, 4, 7, 10, 11, 14, 15, 16, 19, 20, 21}
	expected := []bool{true, false, false, true, false, false, true, false, true, true, false}

	for key, value := range given {
		result := InArray(value, arr)
		assert.Equal(t, expected[key], result)
	}
}

func TestInArrayWithString(t *testing.T) {
	arr := []string{"apple", "banana", "tomato", "grape", "cherry", "mango", "melon"}
	given := []string{"apple", "blueberry", "guava", "cherry", "cranberry", "mango", "melon", "strawberry"}
	expected := []bool{true, false, false, true, false, true, true, false}

	for key, value := range given {
		result := InArray(value, arr)
		assert.Equal(t, expected[key], result)
	}
}
