package w5

import "testing"

func TestSum(t *testing.T) {
	number := []int{1, 2, 3, 4, 5}
	expected := 15
	actual := Sum(number)
}
