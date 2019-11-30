package main

import (
	"fmt"
)

func main() {
	defer printFirst()
	defer printFinish() //ลองเพิ่ม defer
	defer printSecond()

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
