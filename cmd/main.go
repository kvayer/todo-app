package main

import (
	"log"

	todoapp "github.com/kvayer/todo-app"
	"github.com/kvayer/todo-app/pkg/handler"
	"github.com/kvayer/todo-app/pkg/repository"
	"github.com/kvayer/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todoapp.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
