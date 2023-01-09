package main

import (
	"fmt"
)

type Veiculo interface {
	Buzinar()
}

type Carro struct {
	Marca string
	Ano   int
}

type Moto struct {
	Modelo string
	Ano    int
}

func (c *Carro) Buzinar() {
	fmt.Println("beep")
}

func (c *Moto) Buzinar() {
	fmt.Println("boop")
}

func buzina(v Veiculo) {
	v.Buzinar()
}

func main() {

	carro := Carro{
		Marca: "DMC DeLorean",
		Ano:   1982,
	}

	moto := Moto{
		Modelo: "Ciquentinha",
		Ano:    1992,
	}

	veiculos := make([]Veiculo, 2)

	veiculos[0] = &carro
	veiculos[1] = &moto

	for i := 0; i < len(veiculos); i++ {
		buzina(veiculos[i])
	}

	buzina(&carro)
	buzina(&moto)

	carro.Buzinar()
	moto.Buzinar()

}
