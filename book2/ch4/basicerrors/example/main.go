package main

import (
	"book2/ch4/basicerrors"
	"fmt"
)

func main() {
	basicerrors.BasicErrors()
	err := basicerrors.SomeFunc()
	if err != nil {
		fmt.Println("custom error:", err)
	}
}
