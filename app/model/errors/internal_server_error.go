package errors

import (
	"fmt"
	"io"
)

// InternalServerError is custom error of Internal Server Error

type InternalServerError interface {
	Format(s fmt.State, verb rune)
	Error() string
	Unwrap() error
	Cause() error
}

type internalServerError struct {
	error
	*stack
}

func (ise *internalServerError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", ise.Cause())
			ise.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, ise.Error())
	case 'q':
		fmt.Fprintf(s, "%q", ise.Error())
	}
}

func (ise *internalServerError) Error() string {
	return ise.error.Error()
}

func (ise *internalServerError) Unwrap() error {
	return ise.error
}

func (ise *internalServerError) Cause() error {
	return ise.error
}

func (bre *internalServerError) ErrorCode() string {
	return "InternalServerError"
}
