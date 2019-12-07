package main

import "fmt"

func main() {
	name("Norawit")
	math(7, 8)
	nameok("")
	result := math2(50 * 50)
	fmt.Println(result * 500)

}

func name(str string) {
	fmt.Println(str)
}

func math(a int, b int) {
	fmt.Println(a + b)

}

func nameok(str string) {
	fmt.Println("ok")
}

func math2(a int, b int) int {
	output := (a + b)
	return output
}
