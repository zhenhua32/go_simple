package main

import (
	"book2/ch3/currency"
	"fmt"
)

func main() {
	userInput := "15.23"
	pennies, err := currency.ConvertStringDollarsToPennies(userInput)
	if err != nil {
		panic(err)
	}
	fmt.Printf("user input converted to pennies: %d\n", pennies)

	pennies += 15
	dollars := currency.ConvertPenniesToDollarsString(pennies)
	fmt.Printf("after adding 15 pennies: %s\n", dollars)
}
