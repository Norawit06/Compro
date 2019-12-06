package main

import "fmt"

func main() {

	array := []int{1, 2, 3, 4, 5}
	fmt.Println("First Length:", len(array))

	array = append(array, 8)
	fmt.Println("Second Length:", len(array))

}
