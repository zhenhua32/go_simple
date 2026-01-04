package errwrap

import (
	"fmt"

	"github.com/pkg/errors"
)

func WrappedError(e error) error {
	return errors.Wrap(e, "An error occurred in WrappedError")
}

type ErrorTyped struct {
	error
}

func Wrap() {
	e := errors.New("original error")
	fmt.Println("Regular Error -", WrappedError(e))
	fmt.Println("Typed Error -", WrappedError(ErrorTyped{error: errors.New("typed error")}))
	fmt.Println("Nil -", WrappedError(nil))
}
