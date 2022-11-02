package repository

import "rpc-backend-go/app/model"

type UserRepository interface {
	Login(id string, password string) (*model.User, error)
}