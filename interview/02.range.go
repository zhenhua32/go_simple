package main

import "fmt"

func test() {
	s := []int{1, 2, 3}
	m := make(map[int]*int)

	// 这种就坑在使用 *int, 而 v 是整个 range 过程中都是变化的, v 的地址即 &v 就是最后的 3
	for i, v := range s {
		m[i] = &v
	}

	for k, v := range m {
		fmt.Println(k, *v)
	}
}

func test2() {
	s := []int{1, 2, 3}
	m := make(map[int]int)

	for i, v := range s {
		m[i] = v
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func main() {
	test()
	fmt.Println()
	test2()
}
