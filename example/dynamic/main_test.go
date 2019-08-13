package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicMain(t *testing.T) {
	err := InitConfig()
	assert.NoError(t, err)
}
