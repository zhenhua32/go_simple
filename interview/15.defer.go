package main

import "fmt"

func a(n int) (r int) {
	defer func() {
		fmt.Println(r, n)
		r += n
		recover()
	}()

	var f func()
	defer f()

	f = func() {
		r += 2
	}
	return n + 1
}

func main() {
	fmt.Println(a(3))
}
