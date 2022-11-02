package infra

import (
	"fmt"
	"rpc-backend-go/app/model"
	"rpc-backend-go/app/model/errors"
	"rpc-backend-go/app/repository"
	"rpc-backend-go/app/util"
)

type todoRepository struct {
	Conn util.DB
}

func NewTodoRepository(conn util.DB) repository.TodoRepository {
	return &todoRepository{Conn : conn}
}

func (tr *todoRepository) Create(todo model.Todo) (error) {
	fmt.Println("infraのCreate開始")
	if err := tr.Conn.Create(&todo).Error(); err != nil {
		return errors.NewDBError(err)
	}
	return nil
}

func (tr *todoRepository) Update(todoRecord model.Todo, todoRequest model.Todo) (error) {
	fmt.Println("infraのUpdate開始")
	if err := tr.Conn.Update(&todoRecord, &todoRequest).Error(); err != nil {
		return errors.NewDBError(err)
	}
	return nil
}

func (tr *todoRepository) Delete(id int) (error) {
	fmt.Println("infraのDelete開始")
	if err := tr.Conn.Delete(&model.Todo{}).Error(); err != nil {
		return errors.NewDBError(err)
	}
	return nil
}