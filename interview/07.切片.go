package main

import "fmt"

func test() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a[3:3:4]
	fmt.Println(b, len(b), cap(b))
}

func main() {
	test()
}
