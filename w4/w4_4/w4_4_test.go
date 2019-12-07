package main

import "testing"

func Test_math2(t *testing.T) {
	a := math2(7, 9)

	if a != 16 {
		t.Errorf("Want result : 16, but it was %v instead", a)
	}

}
