package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("arquivo.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err != io.EOF {
				fmt.Println(err.Error())
			}
			break
		}
		fmt.Print(string(b))
	}

	// file, err := os.Open("arquivo.txt")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// scaner := bufio.NewScanner(file)
	// for scaner.Scan() {
	// 	fmt.Println(scaner.Text())
	// }
	// err = scaner.Err()

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// file, err := os.ReadFile("arquivo.txt")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("File contents:", string(file))

}
