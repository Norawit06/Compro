package main

import "fmt"

func math(a, b int) int {
	return a + b
}

func main() {
	var num int
	num = math(1, 2)
	fmt.Printf("result = %d", num)

}
