package w7_5

import "fmt"

type Employee struct {
	firstName, lastName string
	age, salary         interface
}

func main() {

	emp1 := Employee{
		firstName: "Sam",
		age:        25,
		salary      1500,
		lastName    "Anderson",
	}
}