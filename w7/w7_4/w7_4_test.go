package w7_4

import "testing"

func TestPlus(t *testing.T) {
	a := plus(1, 2)

	if a != 3 {
		t.Errorf("Want 3, got %v", a)
	}
}

func TestPrism(t *testing.T) {
	a := prism(4, 2, 2)

	if a != 16 {
		t.Errorf("want 16, got %v ", a)
	}
}
