package main

import (
	"log"

	"github.com/puripat-hugeman/go-clean-todo/config"
	"github.com/puripat-hugeman/go-clean-todo/todo/delivery/restful"
	"github.com/puripat-hugeman/go-clean-todo/todo/delivery/server"
	"github.com/puripat-hugeman/go-clean-todo/todo/repository"
	"github.com/puripat-hugeman/go-clean-todo/todo/repository/postgres"
	"github.com/puripat-hugeman/go-clean-todo/todo/usecase"
)

func main() {
	conf, err := config.LoadConfig("config")
	if err != nil {
		log.Fatalln("error: failed to load config:", err.Error())
	}
	db, err := postgres.New(&conf.Postgres)
	if err != nil {
		panic(err)
	}
	repository.Migrate(db)
	todoRepo := repository.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUseCase(todoRepo)
	server := server.NewHttpServer()
	handler := restful.NewHandler(todoUsecase)
	server.SetUpRoutes(*handler)

	if err := server.Serve(conf.Port); err != nil {
		panic(err)
	}

}
