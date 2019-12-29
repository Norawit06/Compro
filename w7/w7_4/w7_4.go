package w7_4

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func prism(length, width, height int) int {
	return length * width * height
}

func init() {
	fmt.Println("Init ...")
}

func main() {
	fmt.Println("WEEK 7 START")
	fmt.Println("-------------------------------")

	a, b := 1, 2
	res := plus(a, b)
	fmt.Println("1 + 2 =", res)

	res = prism(4, 2, 2)
	fmt.Println("Rectangular prism=", res)
	fmt.Println("-------------------------------")
}
