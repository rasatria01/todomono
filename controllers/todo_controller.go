package controllers

import (
	"net/http"
	"sync"
	"todomono/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var (
	todos = []models.Todo{}
	lock  = sync.Mutex{}
)

func CreateTodo(C echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var newTodo models.Todo

	if err := C.Bind(&newTodo); err != nil {
		return err
	}
	newTodo.ID = uuid.New().String()
	todos = append(todos, newTodo)
	return C.JSON(http.StatusCreated, &newTodo)
}

func UpdateTodo(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var updateTodo models.Todo
	if err := c.Bind(&updateTodo); err != nil {
		return err
	}
	id := c.Param("id")
	for i, item := range todos {
		if item.ID == id {
			updateTodo.ID = id
			todos[i] = updateTodo
			return c.JSON(http.StatusOK, updateTodo)
		}
	}
	return c.JSON(http.StatusNotFound, "Todo not Found")
}

func GetTodo(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id := c.Param("id")

	for _, item := range todos {
		if item.ID == id {
			return c.JSON(http.StatusOK, item)
		}
	}
	return c.JSON(http.StatusNotFound, "Todo Not Found")
}

func GetTodos(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	return c.JSON(http.StatusOK, todos)
}

func DeleteTodo(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id := c.Param("id")

	for i, item := range todos {
		if item.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return c.JSON(http.StatusNoContent, nil)
		}
	}
	return c.JSON(http.StatusNotFound, "Todo not Found")
}
