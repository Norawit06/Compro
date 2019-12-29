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
	emp2 := Employee{"Jim", "paul", 29, 9000}

	fmt.Println("Employee 1 ", emp1)
	fmt.Println("Employee 2 ", emp2)
}