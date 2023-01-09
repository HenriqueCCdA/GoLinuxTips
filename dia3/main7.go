package main

import "fmt"

func main() {

	caracteres := []string{"a", "b", "c", "d", "e"}

	// Caracteres é um slice

	caracteres = append(caracteres, "f")

	fmt.Println(caracteres)

	capacidade := 200
	tamanhoDoSile := 10

	caracteres = make([]string, tamanhoDoSile, capacidade)
	caracteres[4] = "4"
	caracteres = append(caracteres, "f")

	fmt.Println(caracteres)

}
