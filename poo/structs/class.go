package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	e := Employee{}
	fmt.Printf("%v", e)

	e.id = 1
	e.name = "Omar"
	fmt.Printf("%v", e)

	e.SetId(4)
	e.SetName("Brayan")
	fmt.Printf("%v", e)

	fmt.Printf(e.GetName())
}
