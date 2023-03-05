package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("Instalando kubctl...!")

	// Install  o kubectl
	// Suportaar linux e macOS
	// Exemplo!!

	command, args := GetKubeCTLInstallCommand()
	cmd := exec.Command(command, args...)

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

}
