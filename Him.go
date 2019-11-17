package main

import (
	"fmt"
)

func main() {
	printFullName("Babel", "Coder")
}

// โปรดสังเกต ชื่อตัวแปรจะมาก่อนชนิดข้อมูล
func printFullName(firstName string, lastName string) {
	fmt.Println(firstName + " " + lastName)
}
