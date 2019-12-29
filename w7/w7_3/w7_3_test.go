package w7_3

import "testing"

func TestInput1ShouldBeDisplay1(t *testing.T) {
	v := Fizzbuzz(1)
	if "2" != v {
		t.Errorf("fizzbuzz of 1 should be '1' but have", v)
	}
}
