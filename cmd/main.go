package main

import (
	"log"

	"github.com/Xploitable/todo-app"
	"github.com/Xploitable/todo-app/package/handler"
	"github.com/Xploitable/todo-app/package/repository"
	"github.com/Xploitable/todo-app/package/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
