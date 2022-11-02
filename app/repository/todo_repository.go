package repository

import "rpc-backend-go/app/model"

type TodoRepository interface {
	Create(model.Todo) (error)
	Update(model.Todo, model.Todo) (error)
	Delete(int) (error)
}