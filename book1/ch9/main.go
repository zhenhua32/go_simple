package main

import (
	"fmt"

	print "tzh.com/example/formatter"
	"tzh.com/example/math"

	"github.com/shopspring/decimal"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)

	amount, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}
	fmt.Println(amount)
}
