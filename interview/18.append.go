package main

import "fmt"

func test() {
	var a = []int{1, 2, 3, 4, 5}
	var r = make([]int, 0)

	for i, v := range a {
		if i == 0 {
			a = append(a, 6, 7)
		}
		r = append(r, v)
	}
	fmt.Println(a)
	fmt.Println(r)
}

func main() {
	test()
}
