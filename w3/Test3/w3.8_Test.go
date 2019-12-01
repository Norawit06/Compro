package main

import "testing"

func TestFib(t *testing.T) {
	a := fib(77)
	if a != 55 {
		t.Errorf("Expect fib is 55 ,but it was %v instead", a)
	}

}
