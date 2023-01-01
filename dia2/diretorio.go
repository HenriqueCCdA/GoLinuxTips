package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func percorreDiretorio(path string, d fs.DirEntry, err error) error {
	fmt.Println("É diretorio? ", d.IsDir())
	info, _ := d.Info()
	fmt.Println("Name: ", info.Name())
	fmt.Println("----------------------\n", info.Name())
	return nil
}

func main() {
	filepath.WalkDir(".", percorreDiretorio)
}
