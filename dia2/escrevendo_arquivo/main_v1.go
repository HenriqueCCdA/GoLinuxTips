package main

import (
	"fmt"
	"os"
)

func main() {
	txt := []byte("Esse é o conteudo do meu arquivo novo")
	file, err := os.OpenFile("arquivo-novo.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.Write(txt)
	file.WriteString("minha string")
	file.WriteAt(txt, 80)
	file.WriteString("\n-----\n")
}
