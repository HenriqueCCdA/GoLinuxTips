package main

import "fmt"

func main() {

	caracteres := [5]string{"a", "b", "c", "d", "e"}

	fmt.Println(len(caracteres))

	for i := 0; i < len(caracteres); i++ {
		fmt.Println(caracteres[i])
	}

	matriz_caracteres := [2][5]string{
		{"a", "b", "c", "d", "e"},
		{"f", "g", "h", "i", "j"},
	}

	for linha := 0; linha < len(matriz_caracteres); linha++ {
		for coluna := 0; coluna < len(matriz_caracteres[linha]); coluna++ {
			fmt.Println(matriz_caracteres[linha][coluna])
		}
	}

}
