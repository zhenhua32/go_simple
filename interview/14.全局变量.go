package main

import "fmt"

var p *int

func a() (*int, error) {
	var i int = 5
	return &i, nil
}

func b() {
	fmt.Println(*p)
}

func test() {
	// var err error
	// p, err = a()
	p, err := a()
	if err != nil {
		fmt.Println(err)
	}
	b()
	fmt.Println(*p)
}

func main() {
	test()
}
