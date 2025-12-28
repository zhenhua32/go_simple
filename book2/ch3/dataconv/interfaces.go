package dataconv

import "fmt"

func CheckType(s any) {
	switch s.(type) {
	case string:
		fmt.Println("is string")
	case int:
		fmt.Println("is int")
	default:
		fmt.Println("unknown type")
	}
}

func Interfaces() {
	CheckType("test")
	CheckType(123)
	CheckType(false)
	var i any
	i = "test"
	if val, ok := i.(string); ok {
		fmt.Println("val is", val)
	}
	if _, ok := i.(int); !ok {
		fmt.Println("un oh! glad we handled this")
	}
}
