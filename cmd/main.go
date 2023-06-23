package main

import (
	"go-todo/internal/delivery/handlers"
	"go-todo/internal/delivery/routes"
	"go-todo/internal/repository"
	"go-todo/internal/repository/database"
	"go-todo/internal/usecase"
)

func main() {
	db := database.NewDB()
	todoRepository := repository.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUsecase(todoRepository)
	todoHandler := handlers.NewTodoHandler(todoUsecase)
	e := routes.NewTodoRouter(todoHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
