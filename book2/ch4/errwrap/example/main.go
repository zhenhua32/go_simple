package main

import (
	"book2/ch4/errwrap"
	"fmt"
)

func main() {
	errwrap.Wrap()
	fmt.Println("-----")

	errwrap.Unwrap()
	fmt.Println("-----")

	errwrap.StackTrace()
}
