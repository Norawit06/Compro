package main

import "fmt"

func math(a, b int) int {
	return a * b
}

func main() {

	num := math(11, 26)
	fmt.Printf("result = %d", num)

}
