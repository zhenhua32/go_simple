package basicerrors

import "fmt"

type CustomError struct {
	Result string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("CustomError occurred with result: %s", c.Result)
}

func SomeFunc() error {
	c := CustomError{Result: "42"}
	return c
}
