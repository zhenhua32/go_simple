package main

import (
	"fmt"
)

func main() {
	text := "Hello, 世界"
	for i, r := range text {
		// fmt.Printf("Character %d: %c (Unicode: %U)\n", i, r, r)
		fmt.Println(i, string(r))
	}

	for i, r := range []rune(text) {
		fmt.Println(i, string(r))
	}
}
