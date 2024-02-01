package moto

import "fmt"

type Motocicleta struct {
	Marca     string
	Model     uint8
	Encendido bool
}

func (m Motocicleta) Encender() {
	if m.Encendido {
		fmt.Println("Ya esta encendida")
		return
	}
}

func (m Motocicleta) Marcha() {
	fmt.Println("CAmbio de marcha!")
}

func (m *Motocicleta) Correr() {
	fmt.Println("Correr!!")
}
