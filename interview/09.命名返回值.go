package main

import "fmt"

func a() (r int) {
	defer func() { r++ }()
	return 0
}

func b() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func c() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main() {
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(c())
}
