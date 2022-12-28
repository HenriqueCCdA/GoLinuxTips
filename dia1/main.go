package main

import "fmt"

// privada, dentro do meu pacote
func dummyFunction(s string) string {
	return fmt.Sprintf("Ran dummy Function. %s\n", s)
}

// publica, dentro do codigo todo
func DummyPublicFuntion(s string) string {
	return dummyFunction(s)
}

func main() {
	// atribuição curta de variavel (declaração + atribuição)
	// short assingment
	s := dummyFunction("'Hello World!")
	fmt.Println(s)

	var u string
	u = dummyFunction("Hellow world again!")
	fmt.Println(u)

	// Controle de fluxo
	// if, else, switch
	var input int
	fmt.Scan(&input)
	if input == 1 {
		fmt.Println("É 1.")
	} else if input == 2 {
		fmt.Println("É 2.")
	} else {
		fmt.Println("Não é 1 ou 2.")
	}

	switch input {
	case 1:
		fmt.Println("É 1.")
	case 2:
		fmt.Println("É 2.")
	default:
		fmt.Println("Não é 1 ou 2.")
	}

	// Estrutura de repetição
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("----")

	// While
	running := true
	for running {
		if input == 3 {
			running = false
		}
	}

	fmt.Println(input)

}
