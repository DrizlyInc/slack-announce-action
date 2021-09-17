package main

import (
	"os"
	"testing"
)

func Test_EnvOrDefault(t *testing.T) {
	varName := "GOLANG_TEST_ENV_VAR"
	defaultValue := "defaultValue"
	realValue := "realValue"

	assertEquals(t, defaultValue, EnvOrDefault(varName, defaultValue))

	os.Setenv(varName, realValue)
	assertEquals(t, realValue, EnvOrDefault(varName, defaultValue))
}

func Test_GetCummulativeStatus(t *testing.T) {
	indicators := []Indicator{
		{Name: "foo", Status: Success},
		{Name: "bar", Status: Success},
	}

	assertEquals(t, Success, GetCummulativeStatus("", indicators))
	assertEquals(t, Failure, GetCummulativeStatus(Failure, indicators))

	indicators[0].Status = Failure
	assertEquals(t, Failure, GetCummulativeStatus("", indicators))
}
