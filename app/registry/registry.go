package registry

import (
	"rpc-backend-go/app/infra"
	"rpc-backend-go/app/interface/api/handler"
	"rpc-backend-go/app/repository"
	"rpc-backend-go/app/util"
)

// SpringBootでいうDIコンテナへの注入
// 最初にserviceとかrepositoryをBeanで宣言するイメージ
// Goでは構造体を定義して、New～で返り値にポインタ指定してインスタンス化することで同じようなことができる

type Registry interface {
	NewUserRepository() repository.UserRepository
	NewUserHandler() handler.UserHandler
}

type registry struct {
	Conn util.DB
	Env *util.Env
}

func NewRegistry(Conn util.DB, Env *util.Env) *registry {
	return &registry{Conn, Env}
}

type appHandler struct {
	user handler.UserHandler
	todo handler.TodoHandler
}

func (ah *appHandler) User() handler.UserHandler {
	return ah.user
}

func (ah *appHandler) Todo() handler.TodoHandler {
	return ah.todo
}

func (r *registry) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{
		user: r.NewUserHandler(),
		todo: r.NewTodoHandler(),
	}
	return appHandler
}

//User
func (r *registry) NewUserRepository() repository.UserRepository {
	return infra.NewUserRepository(r.Conn)
}
func (r *registry) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(r.NewUserRepository())
}

//Todo
func (r *registry) NewTodoRepository() repository.TodoRepository {
	return infra.NewTodoRepository(r.Conn)
}
func (r *registry) NewTodoHandler() handler.TodoHandler {
	return handler.NewTodoHandler(r.NewTodoRepository())
}