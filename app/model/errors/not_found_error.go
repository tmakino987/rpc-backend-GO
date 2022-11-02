package errors

import (
	"fmt"
	"io"
)

// NotFoundError is custom error of Not Found Error
type NotFoundError interface {
	Format(s fmt.State, verb rune)
	Error() string
	Unwrap() error
	Cause() error
}

type notFoundError struct {
	error
	*stack
}

func (nfe *notFoundError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", nfe.Cause())
			nfe.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, nfe.Error())
	case 'q':
		fmt.Fprintf(s, "%q", nfe.Error())
	}
}

func (nfe *notFoundError) Error() string {
	return nfe.error.Error()
}

func (nfe *notFoundError) Unwrap() error {
	return nfe.error
}

func (nfe *notFoundError) Cause() error {
	return nfe.error
}

func (bre *notFoundError) ErrorCode() string {
	return "NotFoundError"
}
