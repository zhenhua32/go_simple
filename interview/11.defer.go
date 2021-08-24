package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func test() {
	a := 1
	b := 2
	defer calc("A", a, calc("10", a, b))
	a = 0
	defer calc("B", a, calc("20", a, b))
	b = 1
}

func main() {
	test()
}
