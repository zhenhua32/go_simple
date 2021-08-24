package main

import "fmt"

func test() {
	const (
		a = iota
		b = iota
	)
	const (
		name = "name"
		c    = iota
		d    = iota
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func main() {
	test()
}
