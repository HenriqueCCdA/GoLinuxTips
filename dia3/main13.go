package main

import "fmt"

func main() {

	alfabeto := []string{"a", "b", "c", "d"}
	fmt.Println("alfabeto:", alfabeto)

	primeirasLetras := alfabeto[0:2]
	fmt.Println(primeirasLetras)

	primeirasLetras[0] = "z"

	primeirasLetras = alfabeto[0:2]
	fmt.Println(primeirasLetras)
	fmt.Println("alfabeto:", alfabeto)

	alfabeto = append(alfabeto, "e")
	primeirasLetras[1] = "y"
	fmt.Println(primeirasLetras)
	fmt.Println("alfabeto:", alfabeto)

}
