package main

import "fmt"

func change(s ...int) {
	s = append(s, 3)
	fmt.Println(s, len(s), cap(s))
}

func test() {
	a := make([]int, 5, 5)
	a[0] = 1
	a[1] = 2

	change(a...)
	fmt.Println(a)

	change(a[:2]...)
	fmt.Println(a)
}

func main() {
	test()
}
