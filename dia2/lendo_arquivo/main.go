package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("TA_PRECO_MEDICAMENTO.CSV")

	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println(err)
		return
	}

}
