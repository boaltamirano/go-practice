package adapter

import "github.com/boaltamirano/go-practice/DesignPatterns/Estructural/Example/bicicleta"

type BicicletaAdapter struct {
	Bicicleta *bicicleta.Bicicleta
}

func (b *BicicletaAdapter) Mover() {
	b.Bicicleta.Avanzar()
}
