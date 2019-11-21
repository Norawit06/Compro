package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Count("apple", "ap"))
	fmt.Println(strings.Count("LOKLKLMLML", "LM"))

	s := "Hi there!"
	fmt.Println("The length of s is", len(s))
	fmt.Println("The first symbol of s is", string(s[0]))
	fmt.Println("The last symbol of s is", string(s[len(s)-1]))
}
