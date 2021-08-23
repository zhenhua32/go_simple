package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	a := "/home//hello/world.txt"
	fmt.Println(filepath.FromSlash(a))
	fmt.Println(filepath.Clean(a))
	fmt.Println(filepath.Abs(a))
}
