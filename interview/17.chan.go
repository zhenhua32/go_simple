package main

import (
	"fmt"
	"runtime"
)

func test() {
	runtime.GOMAXPROCS(1)
	intChan := make(chan int, 1)
	strChan := make(chan string, 1)
	intChan <- 1
	strChan <- "hello"
	select {
	case v := <-intChan:
		fmt.Println(v)
	case v := <-strChan:
		fmt.Println(v)
	}
}

func main() {
	test()
}
