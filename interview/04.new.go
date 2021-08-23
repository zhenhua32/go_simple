package main

import "fmt"

func test() {
	m := new([]int)
	*m = append(*m, 1)
	fmt.Println(*m)
}

func main() {
	test()
}
