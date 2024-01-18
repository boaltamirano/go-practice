package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	// Constructor ejemplo 1
	e := Employee{}
	fmt.Printf("%v\n", e)

	// Constructor ejemplo 2
	e2 := Employee{
		id:       1,
		name:     "Name",
		vacation: true,
	}
	fmt.Printf("%v\n", e2)

	// Constructor ejemplo 1
	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)

	// Constructor ejemplo 4
	e4 := NewEmployee(3, "Omar", true)
	fmt.Printf("%v\n", *e4)
}
