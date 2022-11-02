package infra

import (
	"rpc-backend-go/app/model"
	"rpc-backend-go/app/model/errors"
	"rpc-backend-go/app/repository"
	"rpc-backend-go/app/util"
)

type userRepository struct {
	Conn util.DB
}

func NewUserRepository(conn util.DB) repository.UserRepository {
	return &userRepository{Conn : conn}
}

func (ur *userRepository) Login(mailAddress string, password string) (*model.User, error) {
	user := &model.User{}
	if err := ur.Conn.Raw("SELECT * FROM users WHERE mail_address = ? AND password = ?", mailAddress, password).Scan(user).Error(); err != nil {
		return nil, errors.NewDBError(err)
	}
	return user, nil
}