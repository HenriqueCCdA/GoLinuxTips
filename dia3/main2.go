package main

import "fmt"

func main() {

	var birds int32 = 5
	var bees int16 = 10

	fmt.Println(birds)
	fmt.Println(bees)

	animals := birds + int32(bees)

	fmt.Println((animals))
}
