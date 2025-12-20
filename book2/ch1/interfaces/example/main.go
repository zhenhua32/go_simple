package main

import (
	"book2/ch1/interfaces"
	"bytes"
	"fmt"
)

func main() {
	in := bytes.NewReader([]byte("example"))

	// 这是初始化了一个结构体的指针变量 out
	out := &bytes.Buffer{}
	fmt.Print("stdout on Copy =")
	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}

	fmt.Println("out bytes buffer =", out.String())

	fmt.Print("stdout on PipeExample =")
	if err := interfaces.PipeExample(); err != nil {
		panic(err)
	}
}
