package main

import (
	"fmt"

	print "tzh.com/example/formatter"
	"tzh.com/example/math"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
}
