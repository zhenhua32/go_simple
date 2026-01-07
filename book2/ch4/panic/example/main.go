package main

import (
	"book2/ch4/panic"
	"fmt"
)

func main() {
	fmt.Println("before panic")
	panic.Catcher()
	fmt.Println("after panic")
}
