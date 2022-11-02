package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"rpc-backend-go/app/model"
	"rpc-backend-go/app/model/errors"
	"rpc-backend-go/app/repository"
)

type UserHandler interface {
	Login() echo.HandlerFunc
}

type userHandler struct {
	userRepository repository.UserRepository
}

func NewUserHandler(
	userRepository repository.UserRepository,
) UserHandler {
	return &userHandler{
		userRepository : userRepository,
	}
}

func (uh *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {

		var loginRequest model.LoginRequest
		if err := c.Bind(&loginRequest); err != nil {
			return ResponseFail(c, errors.NewBadRequestError(err))
		}

		fmt.Println("Info of loginRequest: ", loginRequest)

		user, err := uh.userRepository.Login(loginRequest.Mail_address, loginRequest.User_password);
		if err != nil {
			return ResponseFail(c, err)
		}
		if len(user.Id) == 0 {
			return ResponseFail(c, errors.NewNotFoundError("未登録ユーザーです。"))
		}
		fmt.Println("Get user Id: ", user.Id)
		fmt.Println("Get user ID: ", user.ID)
		fmt.Println("Get user Mail_address: ", user.Mail_address)
		fmt.Println("Get user Password: ", user.Password)
		
		return ResponseSuccess(c, WithData(user))
	}
}