package main

import (
	"fmt"
)

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}

	fmt.Println("Employee 3", emp3)

	var emp4 Employee
	fmt.Println("Employee 4", emp4)
}
