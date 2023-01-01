package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	txt := []byte("Esse é o conteudo do meu arquivo novo")
	var perm fs.FileMode
	perm = 0666
	os.WriteFile("arquivo-novo.txt", txt, perm)

	file, err := os.Create("arquivo-novo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.Write(txt)
	file.WriteString("minha string")
	file.WriteAt(txt, 80)
}
