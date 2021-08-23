package main

import "fmt"

func test() {
	// 已经初始化值了
	// 第一个参数定义了当前的长度
	m := make([]int, 3)
	j := append(m, 1, 2)
	fmt.Println(j)

	// 第二个参数定义了底层的容量
	var n = make([]int, 0, 100)
	n = append(n, 1, 2)
	fmt.Println(n)
}

func main() {
	test()
}
