package main

import (
	"fmt"
	"unicode/utf8"
)

func decode() {
	a := []byte("中国")
	fmt.Println(a)
	r, size := utf8.DecodeRune(a)
	fmt.Printf("%c %v %v\n", r, r, size)
}

func encode() {
	a := '中'
	buf := make([]byte, 10)

	n := utf8.EncodeRune(buf, a)
	fmt.Println(buf)
	fmt.Println(n)

}

func length() {
	a := "中国"
	fmt.Println(len(a))
	fmt.Println(utf8.RuneCountInString(a))
}

func main() {
	decode()
	encode()
	length()
}
