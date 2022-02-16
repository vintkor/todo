package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/vintkor/todo"
	"github.com/vintkor/todo/pkg/handler"
	"github.com/vintkor/todo/pkg/repository"
	"github.com/vintkor/todo/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run("5050", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error ocuped while running http server: %s", err)
	}
}
