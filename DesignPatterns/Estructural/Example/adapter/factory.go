package adapter

import (
	"github.com/boaltamirano/go-practice/DesignPatterns/Estructural/Example/auto"
	"github.com/boaltamirano/go-practice/DesignPatterns/Estructural/Example/bicicleta"
	"github.com/boaltamirano/go-practice/DesignPatterns/Estructural/Example/moto"
)

func Factory(s string) Adapter {
	switch s {
	case "auto":
		d := auto.Automovil{}
		return &AutomovilAdapter{&d}
	case "bici":
		d := bicicleta.Bicicleta{}
		return &BicicletaAdapter{&d}
	case "moto":
		d := moto.Motocicleta{}
		return &MotocicletaAdapter{&d}
	}

	return nil
}
