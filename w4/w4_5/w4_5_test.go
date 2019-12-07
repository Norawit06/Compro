package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello Per!"
	actual := hello()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s, got: '%s'", expected, actual)
	}
}
