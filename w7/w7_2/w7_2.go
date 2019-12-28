package w7_2

import (
	"fmt"

	"calhoun.io/testing101"
)

func main() {
	testSum([]int{2, 2, 2, 4}, 10)
	testSum([]int{-1, -2, -3, -4, 5}, -5)
}

func testSum(numbers []int, expected int) {
	sum := testing101.Sum(numbers)
	if sum != expected {
		message := fmt.Sprintf("Expected the sum of %v to be %d but instead got %d!", numbers, expected, sum)
		panic(message)
	}
}
