package main

import "fmt"

func main() {

	func() {
		fmt.Println(("Preciso de uma função separada, pequena e que não vou usar muito"))
	}()

	func(a int) {
		fmt.Println(a)
	}(5)

	f := func(a int) {
		fmt.Println(a)
	}

	f(4)
}
