package main

import "fmt"

type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getName() string {
	return c.name
}

func (c *Computer) getStock() int {
	return c.stock
}

// ToString function
func (c *Computer) String() string {
	s := fmt.Sprintf("Product: %s, with Stock: %d", c.name, c.stock)
	return s
}

type Laptop struct {
	Computer
}

func NewLaptop() IProduct {
	return &Laptop{
		Computer: Computer{
			name:  "Laptop Computer",
			stock: 25,
		},
	}
}

type Desktop struct {
	Computer
}

func NewDesktop() IProduct {
	return &Desktop{
		Computer: Computer{
			name:  "Desktop Computer",
			stock: 12,
		},
	}
}

// This is the factory
func ComputerFactory(computerType IProduct) (IProduct, error) {
	switch computerType.(type) {
	case *Desktop:
		return NewDesktop(), nil
	case *Laptop:
		return NewLaptop(), nil
	default:
		return nil, fmt.Errorf("Unknown computer type: %T", computerType)
	}
}

func main() {
	laptop, _ := ComputerFactory(&Desktop{})
	desktop, _ := ComputerFactory(&Laptop{})

	fmt.Println(laptop)
	fmt.Println(desktop)
}
