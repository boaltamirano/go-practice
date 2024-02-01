// Implementation of design parttern adapter

package main

import (
	"fmt"

	"github.com/boaltamirano/go-practice/DesignPatterns/Estructural/Example/adapter"
)

func main() {

	var t string

	fmt.Print("Digite el tipo de transporte: ")
	fmt.Scanln(&t)

	trasport := adapter.Factory(t)
	trasport.Mover()
}
