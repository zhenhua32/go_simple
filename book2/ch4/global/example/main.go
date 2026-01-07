package main

import (
	"book2/ch4/global"
)

func main() {
	if err := global.UseLog(); err != nil {
		panic(err)
	}
}
