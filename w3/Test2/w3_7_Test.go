package main

import (
	"testing"
)

func TestInput1ShouldBeDisplay1(t *testing.T) {

	v := FizzBuzz(1)

	if "1" != v {
		t.Error("fizzbuzz of 1 should be '1' but have", v)
	}
}
