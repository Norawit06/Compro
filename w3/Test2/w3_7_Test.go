package main

import "testing"

func testInput1ShouldBeDisplay1(t *testing.T) {

	v := FizzBuzz(1)

	if "1" != v {
		t.Error("Fizzbuzz of 1 Should be `1` but have", v)
	}
}
