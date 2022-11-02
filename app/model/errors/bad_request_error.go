package errors

import (
	"fmt"
	"io"
)

// BadRequestError is custom error of Bad Request Error
type BadRequestError interface {
	Format(s fmt.State, verb rune)
	Error() string
	Unwrap() error
	Cause() error
}

type badRequestError struct {
	error
	*stack
}

func (bre *badRequestError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", bre.Cause())
			bre.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, bre.Error())
	case 'q':
		fmt.Fprintf(s, "%q", bre.Error())
	}
}

func (bre *badRequestError) Error() string {
	return bre.error.Error()
}

func (bre *badRequestError) Unwrap() error {
	return bre.error
}

func (bre *badRequestError) Cause() error {
	return bre.error
}

func (bre *badRequestError) ErrorCode() string {
	return "BadRequest"
}
