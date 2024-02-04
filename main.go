package main

import (
	"net/http"
	"todomono/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.GET("/todos", controllers.GetTodos)
	e.POST("/todos", controllers.CreateTodo)
	e.GET("/todos/:id", controllers.GetTodo)
	e.PUT("/todos/:id", controllers.UpdateTodo)
	e.DELETE("/todos/:id", controllers.DeleteTodo)

	e.Logger.Fatal(e.Start(":8080"))
}
