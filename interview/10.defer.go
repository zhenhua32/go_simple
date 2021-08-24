package main

import "fmt"

type Person struct {
	age int
}

func test() {
	p := &Person{age: 28}
	defer fmt.Println(p.age)
	defer func(a *Person) {
		fmt.Println(a.age)
	}(p)
	defer func() {
		fmt.Println(p.age)
	}()

	p.age = 30
}

func main() {
	test()
}
