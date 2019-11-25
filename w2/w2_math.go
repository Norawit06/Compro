package main

import "fmt"

func math(a, b float64) float64 {
	return a + b
}

func main() {

	num := math(11.215151, 65.2151515214)
	fmt.Printf("result = %.30f", num)

}
