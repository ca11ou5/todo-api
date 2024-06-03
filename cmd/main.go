package main

import (
	env "github.com/ilyakaznacheev/cleanenv"
	"log"
	"todo-list/configs"
	"todo-list/internal/handler"
	"todo-list/internal/handler/http"
	"todo-list/internal/repository"
	"todo-list/internal/service"
)

//	@title			Todo List API
//	@version		1.0
//	@description	This is a sample server for a todo list.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/

func main() {
	var cfg configs.Config

	err := env.ReadConfig("./envs/local.env", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(&cfg)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	http.StartListening(&cfg, h)
}
