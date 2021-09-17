package main

import "testing"

func assertEquals(t *testing.T, expected, actual string) {
	if actual != expected {
		t.Fatalf("Expected %s to equal %s", actual, expected)
	}
}
