package main


import "fmt"

func main(){
	var name string
	var age int
	n, err :=fmt.Sscanf("Earth is 25 years old", "%s is %d years old", &name, &age)
	