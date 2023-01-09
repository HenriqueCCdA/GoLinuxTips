package main

import "fmt"

type Carro struct {
	Marca string // atributos
	Ano   int
}

func main() {

	// Struct anônima
	carro := struct {
		Marca string
		Ano   int
	}{
		Marca: "DMC Delorean",
		Ano:   1982,
	}

	fmt.Println(carro)
	fmt.Println(carro.Marca)
	fmt.Println(carro.Ano)

	carro1 := Carro{
		Marca: "DMC Delorean",
		Ano:   1982,
	}

	fmt.Println(carro1)
	fmt.Println(carro1.Marca)
	fmt.Println(carro1.Ano)

}
