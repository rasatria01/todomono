package main

import (
	"net/http"
	"os"
	"todomono/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func main() {
	e := echo.New()
	logger := zerolog.New(os.Stdout)
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:     true,
		LogStatus:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", v.URI).
				Int("Status", v.Status).
				Int("Duration", int(v.Latency)).
				Msg("request")

			return nil
		},
	}))
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
