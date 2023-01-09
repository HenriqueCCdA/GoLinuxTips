package main

import "fmt"

func main() {

	// Conteudo: 0 0 0 0 0 0 1 0 | 0 |      0     |       8     |
	// Enderoço: 0 1 2 3 4 5 6 7 | 8 | 9 10 11 12 | 13 14 15 16 | 17 18 19 20
	// Variavel: ---- x -------- | b | --- px --- | ---- pb --- |

	var x int
	x = 10 // 8 bytes
	var b bool
	b = false // 1 byte

	pointerOfX := &x
	pointerOfB := &b

	fmt.Println(x)
	fmt.Println(b)
	fmt.Println(pointerOfX)
	fmt.Println(pointerOfB)

}
