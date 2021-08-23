package main

import (
	"fmt"
)

func Test1() {
	defer fmt.Println("test1")
	defer fmt.Println("test2")
	defer fmt.Println("test3")
	fmt.Println("hello")
}

func main() {
	Test1()
}
