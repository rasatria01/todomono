package controllers

import (
	"net/http"
	"sync"
	"todomono/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var (
	todo = []models.Todo{}
	seq  = 1
	lock = sync.Mutex{}
)

func CreateTodo(C echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var newTodo models.Todo

	if err := C.Bind(&newTodo); err != nil {
		return err
	}
	newTodo.ID = uuid.New().String()
	todo = append(todo, newTodo)
	return C.JSON(http.StatusCreated, &newTodo)
}
