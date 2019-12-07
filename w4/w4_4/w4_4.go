package main

import "fmt"

func main() {
	name("Norawit")
	math(7, 8)
	nameok("")

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
