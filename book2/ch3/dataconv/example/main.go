package main

import (
	"book2/ch3/dataconv"
	"fmt"
)

func main() {
	dataconv.ShowConv()
	fmt.Println("----")
	if err := dataconv.Strconv(); err != nil {
		panic(err)
	}
	fmt.Println("----")
	dataconv.Interfaces()
}
