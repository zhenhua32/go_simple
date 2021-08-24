package main

import "fmt"

func test() {
	m := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func main() {
	test()
}
