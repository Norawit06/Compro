package main

import (
	"fmt"
)

func main() {
	printFullName("Babel", "Coder")
}

func printFullName(firstName string, lastName string) {
	fmt.Println(firstName + " " + lastName)
}
