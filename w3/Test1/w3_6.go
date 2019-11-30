package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func prism(length, width, height int) int {
	return length * width * height
}

func init() {
	fmt.Println("init ...")
}

func main() {
	fmt.Println("Week 4 Start ...")
	fmt.Println("-----------------------------")

	a, b := 1, 2
	res := plus(a, b)
	fmt.Println("1 + 2 =", res)

	res = prism(4, 2, 2)
	fmt.Println("Rectangular prism=", res)

	fmt.Println("------------------------------")

}
