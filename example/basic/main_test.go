package main

import (
	"testing"
)

func TestBasicMain(t *testing.T) {
	err := InitConfig()
	if err != nil {
		t.Error(err)
	}
}
