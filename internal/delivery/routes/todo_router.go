package routes

import (
	"go-todo/internal/delivery/handlers"

	"github.com/labstack/echo/v4"
)

func NewTodoRouter(th handlers.TodoHandler) *echo.Echo {
	e := echo.New()
	t := e.Group("/todos")
	t.GET("", th.GetAllTodos)
	t.GET("/:todoId", th.GetTodoById)
	t.POST("", th.CreateTodo)
	t.PUT("/:todoId", th.UpdateTodo)
	t.DELETE("/:todoId", th.DeleteTodo)

	return e
}
