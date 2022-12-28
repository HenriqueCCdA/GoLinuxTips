package main

import "fmt"

func soma(a, b int) int {
	return a + b
}

func dividir(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func subtracao(a, b int) int {
	return a - b
}

func mutiplicar(a, b int) int {
	return a * b
}

func main() {
	// TODO: Pedir os numeros par o usuário?

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			fmt.Println(i)
			fmt.Println(j)

			resposta := soma(i, j)
			fmt.Printf("Soma: %d\n", resposta)

			resposta = dividir(i, j)
			fmt.Printf("Divisão: %d\n", resposta)

			resposta = subtracao(i, j)
			fmt.Printf("Subtração: %d\n", resposta)

			resposta = mutiplicar(i, j)
			fmt.Printf("Multiplicação: %d", resposta)
		}

	}
}
