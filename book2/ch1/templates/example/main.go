package main

import (
	"book2/ch1/templates"
	"fmt"
)

func main() {
	if err := templates.RunTemplate(); err != nil {
		panic(err)
	}
	fmt.Println("--------------")

	if err := templates.InitTemplates(); err != nil {
		panic(err)
	}
	fmt.Println("--------------")

	if err := templates.HTMLDifferences(); err != nil {
		panic(err)
	}
	fmt.Println("--------------")

}
