package main

import "fmt"

func main() {
	var i, j string
	var p *string
	i = "Earth"
	j = "Peach"

	fmt.Printf("i = %c\n", i)
	fmt.Printf("address of i %v\n", &i)
	fmt.Printf("address of j %v\n", &j)
	p = &i
	fmt.Printf("p = %v\n", p)
	fmt.Printf("*p = %v\n", *p)
	p = &j
	fmt.Printf("p = %v\n", p)
	fmt.Printf("*p = %v\n", *p)

}
