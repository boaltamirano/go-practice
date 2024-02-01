package adapter

import "github.com/boaltamirano/go-practice/DesignPatterns/Estructural/Example/moto"

type MotocicletaAdapter struct {
	Motocic *moto.Motocicleta
}

func (m *MotocicletaAdapter) Mover() {
	m.Motocic.Encender()
	m.Motocic.Marcha()
	m.Motocic.Correr()
}
