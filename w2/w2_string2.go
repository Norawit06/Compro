package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Pretty Earth", "earth"))
	fmt.Println(strings.Contains("Pretty Earth", "Earth"))

	fmt.Println(strings.Count("Nice Earth", "e"))
	fmt.Println(strings.Count("Nice Earth", "h"))
	fmt.Println(strings.Count("Nice Earth", ""))

}
