package main

import (
	"fmt"
	"math"
)

func main() {
	var a int8
	var b int16
	var c int32
	var d int64

	a = math.MaxInt8
	b = math.MaxInt16
	c = math.MaxInt32
	d = math.MaxInt64

	fmt.Println("Tamanho máximo int8 :: ", a)
	fmt.Println("Tamanho máximo int16 :: ", b)
	fmt.Println("Tamanho máximo int32 :: ", c)
	fmt.Println("Tamanho máximo int64 :: ", d)

	a = math.MinInt8
	b = math.MinInt16
	c = math.MinInt32
	d = math.MinInt64

	fmt.Println("Tamanho minimo int8 :: ", a)
	fmt.Println("Tamanho minimo int16 :: ", b)
	fmt.Println("Tamanho minimo int32 :: ", c)
	fmt.Println("Tamanho minimo int64 :: ", d)
}
