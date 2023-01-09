package main

import (
	"encoding/json"
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

	veiculos := []Veiculo{&carro, &moto}

	// JSON

	data, err := json.Marshal(veiculos)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(data))

	s := "[{\"Marca\":\"DMC DeLorean\",\"Ano\":1982}]"

	carros := []Carro{}

	err = json.Unmarshal([]byte(s), &carros)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(carros)
}
