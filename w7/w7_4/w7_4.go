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
