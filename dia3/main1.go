package main

import "fmt"

func main() {

	var i uint64 = 18446744073709551615

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		i++
		fmt.Println(i)
	}
}
