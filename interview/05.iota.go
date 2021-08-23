package main

import "fmt"

func test() {
	const (
		a = iota
		_
		b
		c = "a"
		d
		e = iota
		f
	)
	fmt.Println(a, b, c, d, e, f)
}

func main() {
	test()
}
