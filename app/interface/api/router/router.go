package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"rpc-backend-go/app/interface/api/handler"
)

func StartRouter(e *echo.Echo, appHandler handler.AppHandler) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*", "http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))


	InitRouting(e, appHandler)
	return e
}

func InitRouting(e *echo.Echo, appHandler handler.AppHandler) {

	con := e.Group("/api")
	{
		con.GET("/healthcheck", echo.HandlerFunc(func(c echo.Context) error { return c.NoContent(http.StatusOK) }))
		con.OPTIONS("/*", echo.HandlerFunc(func(c echo.Context) error { return c.NoContent(http.StatusOK) }))
		con.POST("/auth/login", appHandler.User().Login())
		con.POST("/todo/register", appHandler.Todo().Create())
		con.POST("/todo/update", appHandler.Todo().Update())
		con.GET("/todo/delete/:id", appHandler.Todo().Delete())
	}
}
