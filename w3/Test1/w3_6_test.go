package main

import (
	"testing"
)

func TestPlus(t *testing.T) {
	a := plus(1, 8)

	if a != 3 {
		t.Errorf("Expected result of 3, but it was  %v instead ", a)
	}
}

func TestPrism(t *testing.T) {
	a := prism(4, 2, 2)
	if a != 28 {
		t.Errorf("Expected result of 3, but it was  %v instead ", a)
	}
}
