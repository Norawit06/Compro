package main

import "fmt"

func main() {
	fmt.Print("input : ")
	var text string
	var age int
	fmt.Scan(&text)
	fmt.Println(`read "`, text, `"from input`)
	fmt.Println("input your age :")
	fmt.Scan(&age)
	fmt.Printf("your age :%d ", age)

}
