package adapter

import "github.com/boaltamirano/go-practice/DesignPatterns/Estructural/Example/auto"

type AutomovilAdapter struct {
	Automovil *auto.Automovil
}

func (a *AutomovilAdapter) Mover() {
	a.Automovil.Encender()
	a.Automovil.Acelerar()
}
