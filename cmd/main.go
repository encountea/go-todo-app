package main

import (
	"log"

	"github.com/encountea/todo-app"
	"github.com/encountea/todo-app/pkg/handler"
	"github.com/encountea/todo-app/pkg/repository"
	"github.com/encountea/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	app := new(todo.Server)

	log.Println("Server is running")
	if err := app.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error to connect to server: %s", err.Error())
	}
}