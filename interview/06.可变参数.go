package main

import "fmt"

func a(num ...int) {
	num[0] = 18
}

func test() {
	i := []int{1, 2, 3}
	a(i...)
	fmt.Println(i)
}

func main() {
	test()
}
