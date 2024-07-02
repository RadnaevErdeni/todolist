package main

import (
	todo "TODO"
	"TODO/handler"
	"TODO/repository"
	"TODO/service"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := repository.DBC(repository.Conf{
		Host:     "localhost",
		Port:     "8080",
		Username: "postgres",
		Password: "PostgreSQL",
		BDname:   "todo",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("Failed to initialize db:%s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("localhost", "8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error running http server: %s", err.Error())
	}
}
