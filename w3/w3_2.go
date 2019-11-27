package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Printf("fib(10)=%d\n", fib(10))
	fmt.Printf("fib(10)=%d\n", fib(11))
	fmt.Printf("fib(10)=%d\n", fib(12))

}
