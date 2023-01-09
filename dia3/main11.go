package main

import "fmt"

func soma(num int, valores ...int) (int, int, error) {

	if num == 0 {
		return 0, 0, fmt.Errorf("Num não pode ser zero")
	}

	for _, valor := range valores {
		num += valor
	}
	return num, len(valores), nil
}

func main() {
	fmt.Println(soma(0))
	fmt.Println(soma(10))
	fmt.Println(soma(10, 5))
	fmt.Println(soma(10, 5, 5, 5, 5))
}
