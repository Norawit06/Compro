package main

import "fmt"

func main() {
	var i int
	var p *int

	i = 10

	fmt.Printf("i = %c\n", i)
	fmt.Printf("address of i %v\n", &i)
	p = &i
	fmt.Printf("p = %v\n", p)
	fmt.Printf("*p = %v\n", *p)

}
