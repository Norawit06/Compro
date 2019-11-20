package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Pretty Earth", "earth"))
	fmt.Println(strings.Contains("Pretty Earth", "Earth"))

}
