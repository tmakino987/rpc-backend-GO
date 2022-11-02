package main

import (
	"rpc-backend-go/app/interface/api/router"

	"github.com/labstack/echo/v4"

	"rpc-backend-go/app/registry"
	"rpc-backend-go/app/util"
)

func main() {
	env := util.NewEnv()

	registry := registry.NewRegistry(
		util.NewDB(env.GetEnvValue("DSN")),
		env,
	)
	handler := registry.NewAppHandler()

	e := echo.New()
	router.StartRouter(e, handler)
	e.Logger.Fatal(e.Start(":8080"))
}
