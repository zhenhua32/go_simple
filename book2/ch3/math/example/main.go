package main

import (
	"book2/ch3/math"
	"fmt"
)

func main() {
	math.Examples()
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", math.Fib(i))
	}
	fmt.Println()
}
