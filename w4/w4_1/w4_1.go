package main

import "fmt"

func main() {
	var name [5]string

	name[0] = "pine"
	name[1] = "app"
	name[2] = "ple"
	name[3] = "pen"
	name[4] = "cil"

	fmt.Println(name[0])

	names := [3]string{"pine", "app", "ple"}
	_ = names
	fmt.Println(name[0], name[1], name[2])
}
