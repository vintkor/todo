package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/vintkor/todo"
	"github.com/vintkor/todo/pkg/handler"
	"github.com/vintkor/todo/pkg/repository"
	"github.com/vintkor/todo/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("Cant read config")
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error ocuped while running http server: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
