package main

import "fmt"

func main() {
	alfabeto := make(map[string]string)
	alfabeto["vogais"] = "aeiou"
	alfabeto["consoantes"] = "bcdfg...."

	fmt.Println(alfabeto)

	alfabeto = map[string]string{
		"vogais":     "aeiou",
		"consoantes": "bcdfg...",
	}
	fmt.Println(alfabeto)

	fmt.Println(alfabeto["vogais"])

}
