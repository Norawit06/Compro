package w5

import "testing"

func TestSum(t *testing.T) {
	number := []int{1, 2, 3, 4, 5}
	expected := 15
	actual := Sum(number)

	if actual != expected{
		t.Errorf("expected the sum of %v to be %d but instead got %d!", number,expected,actual
	}
}
