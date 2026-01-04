package errwrap

import (
	"fmt"

	"github.com/pkg/errors"
)

func Unwrap() {
	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")
	fmt.Println("wrapped error:", err)

	switch errors.Cause(err).(type) {
	case ErrorTyped:
		fmt.Println("This is a typed error:", err)
	default:
		fmt.Println("This is a unknown error:", err)
	}
}

func StackTrace() {
	err := error(ErrorTyped{error: errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")
	fmt.Printf("%+v\n", err)
}
