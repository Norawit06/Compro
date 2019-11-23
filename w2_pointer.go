package main

import "fmt"

func main() {
	var i string
	var p *string

	i = "Earth"

	fmt.Printf("i = %c\n", i)
	fmt.Printf("address of i %v\n", &i)
	p = &i
	fmt.Printf("p = %v\n", p)
	fmt.Printf("*p = %v\n", *p)

}
