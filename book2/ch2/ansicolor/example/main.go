package main

import (
	"book2/ch2/ansicolor"
	"fmt"
)

func main() {
	r := ansicolor.ColorText{
		TextColor: ansicolor.Red,
		Text:      "hello world",
	}
	fmt.Println(r.String())
	r.TextColor = ansicolor.Green
	fmt.Println(r.String())
	r.TextColor = ansicolor.ColorNone
	fmt.Println(r.String())
	fmt.Println(r.Text)
}
