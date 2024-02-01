package adapter

import (
	"github.com/boaltamirano/go-practice/DesignPatterns/Estructural/Example/auto"
	"github.com/boaltamirano/go-practice/DesignPatterns/Estructural/Example/bicicleta"
)

func Factory(s string) Adapter {
	switch s {
	case "auto":
		d := auto.Automovil{}
		return &AutomovilAdapter{&d}
	case "bici":
		d := bicicleta.Bicicleta{}
		return &BicicletaAdapter{&d}
	}

	return nil
}
