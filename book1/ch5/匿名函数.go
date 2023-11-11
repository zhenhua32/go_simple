package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 定义一个匿名函数, 并立即调用
	func() {
		fmt.Println("Hello, world!", rand.Intn(10))
	}()
}
