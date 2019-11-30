package main

import (
	"fmt"
)

func main() {
	printFirst()
	printFinish()
	printSecond()

}

func printFirst() {
	fmt.Println("First")
}

func printSecond() {
	fmt.Println("Second")
}

func printFinish() {
	fmt.Println("Close")
}
