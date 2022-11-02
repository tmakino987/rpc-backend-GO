package errors

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

type ErrorCode interface {
	ErrorCode() string
}

const InternalServerErrorMessage = "internal server error"

func New(msg string) error {
	return errors.New(msg)
}

func NewDBError(err error) error {
	switch err {
	case gorm.ErrRecordNotFound:
		return NewNotFoundError(err)
	case gorm.ErrInvalidData, gorm.ErrInvalidField, gorm.ErrInvalidValue, gorm.ErrUnsupportedRelation, gorm.ErrPrimaryKeyRequired:
		return NewBadRequestError(err)
	default:
		log.Errorf("%v", err)
		return NewInternalServerError(InternalServerErrorMessage)
	}
}

func NewNotFoundError(value interface{}) error {
	switch typedValue := value.(type) {
	case error:
		return &notFoundError{errors.WithMessage(typedValue, "Not Found Error"), callers(typedValue)}
	default:
		return &notFoundError{fmt.Errorf("Not Found Error: %v", typedValue), callers(nil)}
	}
}

func NewBadRequestError(value interface{}) error {
	switch typedValue := value.(type) {
	case error:
		return &badRequestError{errors.WithMessage(typedValue, "Bad Request Error"), callers(typedValue)}
	default:
		return &badRequestError{fmt.Errorf("Bad Request Error: %v", typedValue), callers(nil)}
	}
}

func NewInternalServerError(value interface{}) error {
	switch typedValue := value.(type) {
	case error:
		return &internalServerError{errors.WithMessage(typedValue, "Internal Server Error"), callers(typedValue)}
	default:
		return &internalServerError{fmt.Errorf("Internal Server Error: %v", typedValue), callers(nil)}
	}
}
