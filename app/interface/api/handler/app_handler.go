package handler

import (
	"fmt"
	"net/http"
	"rpc-backend-go/app/model/errors"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type AppHandler interface {
	User() UserHandler
	Todo() TodoHandler
}

func ResponseFail(c echo.Context, err error) error {
	fmt.Printf("error: %+v", err)

	log.Errorf("%v", err)

	var statusCode int
	if e, ok := err.(errors.ErrorCode); ok {
		switch e.ErrorCode() {
		case "BadRequest":
			statusCode = http.StatusBadRequest
		case "NotFoundError":
			statusCode = http.StatusNotFound
		case "InternalServerError":
			statusCode = http.StatusInternalServerError
		}
	} else {
		statusCode = http.StatusInternalServerError
	}

	return c.JSON(statusCode, struct {
		Code int `json:"code"`
		Message string `json:"message"`
	}{
		Code: statusCode,
		Message: err.Error(),
	})
}

type responseOption struct {
	message string
	data    interface{}
}

type Option func(*responseOption)

func WithMessage(message string) Option {
	return func(ro *responseOption) {
		ro.message = message
	}
}

func WithData(data interface{}) Option {
	return func(ro *responseOption) {
		ro.data = data
	}
}

func ResponseSuccess(c echo.Context, opts ...Option) error {
	responseOption := &responseOption{}
	for _, opt := range opts {
		opt(responseOption)
	}

	var message string
	if len(responseOption.message) > 0 {
		message = responseOption.message
	} else {
		message = "Success"
	}
	var header = c.Response().Header()
	header.Set("Cache-Control", "no-store")
	header.Add("Pragma", "no-cache")

	if responseOption.data != nil {
		return c.JSON(http.StatusOK, struct {
			Code int `json:"code"`
			Message string `json:"message"`
			Data interface{} `json:"data"`
		}{
			Code : http.StatusOK,
			Message : message,
			Data : responseOption.data,
		})
	}

	return c.JSON(http.StatusOK, struct {
		Code int `json:"code"`
		Message string `json:"message"`
	}{
		Code : http.StatusOK,
		Message : message,
	})
}