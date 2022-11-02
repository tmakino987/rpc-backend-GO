package handler

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"

	"rpc-backend-go/app/model"
	"rpc-backend-go/app/model/errors"
	"rpc-backend-go/app/repository"
)

type TodoHandler interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type todoHandler struct {
	todoRepository repository.TodoRepository
}

func NewTodoHandler(
	todoRepository repository.TodoRepository,
) TodoHandler {
	return &todoHandler{
		todoRepository : todoRepository,
	}
}

func (th *todoHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		fmt.Println("Createメソッド開始")
		
		var todoRequest model.TodoRequest
		if err := c.Bind(&todoRequest); err != nil {
			return ResponseFail(c, errors.NewBadRequestError(err))
		}

		todoEntity := model.Todo{
			User_id: todoRequest.User_id,
			Status: todoRequest.Status,
			Title: todoRequest.Title,
			Memo: todoRequest.Memo,
			Reminder_date: todoRequest.Reminder_date,
			Deadline_date: todoRequest.Deadline_date,
			Repeat_type: todoRequest.Repeat_type,
			Genre: todoRequest.Genre,
			Important: todoRequest.Important,
			Priority: todoRequest.Priority,
			Image_file_path: todoRequest.Image_file_path,
		}

		fmt.Println("Createにデータ投入")
		err := th.todoRepository.Create(todoEntity);
		if err != nil {
			return ResponseFail(c, err)
		}

		return ResponseSuccess(c)
	}
}

func (th *todoHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		fmt.Println("Updateメソッド開始")
		
		var todoRequest model.TodoRequest
		if err := c.Bind(&todoRequest); err != nil {
			return ResponseFail(c, errors.NewBadRequestError(err))
		}

		if todoRequest.Id == 0 {
			return ResponseFail(c, errors.NewBadRequestError("更新対象のIDが未指定"))
		}

		todoRecord := model.Todo{
			Id: todoRequest.Id,
		}

		todoValue := model.Todo{
			User_id: todoRequest.User_id,
			Status: todoRequest.Status,
			Title: todoRequest.Title,
			Memo: todoRequest.Memo,
			Reminder_date: todoRequest.Reminder_date,
			Deadline_date: todoRequest.Deadline_date,
			Repeat_type: todoRequest.Repeat_type,
			Genre: todoRequest.Genre,
			Important: todoRequest.Important,
			Priority: todoRequest.Priority,
			Image_file_path: todoRequest.Image_file_path,
		}

		err := th.todoRepository.Update(todoRecord, todoValue);
		if err != nil {
			return ResponseFail(c, err)
		}

		return ResponseSuccess(c)
	}
}

func (th *todoHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {

		fmt.Println("Deleteメソッド開始")
		
		var todoId int
		var err error
		id := c.Param("id")
		todoId, err = strconv.Atoi(id)
		if err != nil {
			return ResponseFail(c, err)
		}

		err = th.todoRepository.Delete(todoId);
		if err != nil {
			return ResponseFail(c, err)
		}

		return ResponseSuccess(c)
	}
}
