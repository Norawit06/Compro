package main

import "fmt"

func main() {
	fmt.Print("input Name : ")
	var text string
	var age int
	fmt.Scan(&text)
	fmt.Println(`Name`, text)
	fmt.Println("input your age :")
	fmt.Scan(&age)
	fmt.Printf("your age :%d ", age)

}
