package main

import "fmt"

func a() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func b() (i int) {
	defer func() {
		i++
	}()
	return i
}

func main() {
	fmt.Println(a())
	fmt.Println()
	fmt.Println(b())
}
