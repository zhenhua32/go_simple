package main

import (
	"book2/ch3/nulls"
	"fmt"
)

func main() {
	if err := nulls.BaseEncoding(); err != nil {
		panic(err)
	}
	fmt.Println("----")

	if err := nulls.PointerEncoding(); err != nil {
		panic(err)
	}
	fmt.Println("----")

	if err := nulls.NullEncoding(); err != nil {
		panic(err)
	}
	fmt.Println("----")
}
