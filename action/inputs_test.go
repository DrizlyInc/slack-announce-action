package main

import (
	"os"
	"testing"
)

func Test_EnvOrDefault(t *testing.T) {
	varName := "GOLAG_TEST_ENV_VAR"
	defaultValue := "defaultValue"
	realValue := "realValue"

	assertEquals(t, defaultValue, EnvOrDefault(varName, defaultValue))

	os.Setenv(varName, realValue)
	assertEquals(t, realValue, EnvOrDefault(varName, defaultValue))
}
