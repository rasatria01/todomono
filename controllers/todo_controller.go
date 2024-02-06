package controllers

import (
	"net/http"
	"sync"
	Config "todomono/config"
	"todomono/models"

	"github.com/labstack/echo/v4"
)

var (
	todos = []models.Todo{}
	lock  = sync.Mutex{}
)

func CreateTodo(C echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	db := Config.GetDB()

	newTodo := new(models.Todo)

	if err := C.Bind(newTodo); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return C.JSON(http.StatusBadRequest, data)
	}

	if err := db.Create(&newTodo).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return C.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Create Successfully",
		"data":    newTodo,
	}
	return C.JSON(http.StatusOK, response)

	// newTodo.ID = uuid.New().String()
	// todos = append(todos, newTodo)
	// return C.JSON(http.StatusCreated, &newTodo)
}

func UpdateTodo(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var updateTodo models.Todo
	if err := c.Bind(&updateTodo); err != nil {
		return err
	}
	// id := c.Param("id")
	// for i, item := range todos {
	// 	if item.ID == id {
	// 		updateTodo.ID = id
	// 		todos[i] = updateTodo
	// 		return c.JSON(http.StatusOK, updateTodo)
	// 	}
	// }
	return c.JSON(http.StatusNotFound, "Todo not Found")
}

func GetTodo(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	db := Config.GetDB()

	id := c.Param("id")

	var todo = models.Todo{}
	if err := db.First(&todo, id); err.Error != nil {
		data := map[string]interface{}{
			"message": err.Error.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	response := map[string]interface{}{
		"message": "data fetch succesfully",
		"data":    todo,
	}
	return c.JSON(http.StatusOK, response)
}

func GetTodos(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	db := Config.GetDB()

	var todo []*models.Todo

	if err := db.Find(&todo); err.Error != nil {
		data := map[string]interface{}{
			"message": err.Error.Error(),
		}
		return c.JSON(http.StatusOK, data)
	}

	response := map[string]interface{}{
		"message": "data fetch succesfully",
		"data":    todo,
	}
	return c.JSON(http.StatusOK, response)
}

func DeleteTodo(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	db := Config.GetDB()

	id := c.Param("id")
	todo := new(models.Todo)
	if err := db.Delete(&todo, id); err.Error != nil {
		data := map[string]interface{}{
			"message": err.Error.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	// for i, item := range todos {
	// 	if item.ID == id {
	// 		todos = append(todos[:i], todos[i+1:]...)
	// 		return c.JSON(http.StatusNoContent, nil)
	// 	}
	// }
	response := map[string]interface{}{
		"message": "Delete Successfully!",
	}
	return c.JSON(http.StatusOK, response)
}
