package errors

import (
	"fmt"
	"runtime"

	"github.com/pkg/errors"
)

type stack []uintptr

func (s *stack) StackTrace() errors.StackTrace {
	f := make([]errors.Frame, len(*s))
	for i := 0; i < len(f); i++ {
		f[i] = errors.Frame((*s)[i])
	}
	return f
}

func (s *stack) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case st.Flag('+'):
			for _, pc := range *s {
				f := errors.Frame(pc)
				fmt.Fprintf(st, "\n%+v", f)
			}
		}
	}
}

func callers(err error) *stack {
	if e, ok := err.(stackTracer); ok {
		var originStack []uintptr
		for _, f := range e.StackTrace() {
			originStack = append(originStack, uintptr(f))
		}

		var stack stack = append(originStack, *newStack()...)
		return &stack
	}
	return newStack()
}

func newStack() *stack {
	const depth = 10
	const skip = 3
	var pcs [depth]uintptr
	n := runtime.Callers(skip, pcs[:])
	var st stack = pcs[0:n]
	return &st
}
