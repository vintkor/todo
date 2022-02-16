package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/vintkor/todo"
)

func main() {
	server := new(todo.Server)
	if err := server.Run("5050"); err != nil {
		log.Fatalf("Error ocuped while running http server: %s", err)
	}
}
