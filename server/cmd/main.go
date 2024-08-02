package main

import (
	"log"
	"net/http"

	"github.com/phzeng0726/go-graphql-server-template/internal/config"
	"github.com/phzeng0726/go-graphql-server-template/internal/database"
	delivery "github.com/phzeng0726/go-graphql-server-template/internal/delivery/http"
	"github.com/phzeng0726/go-graphql-server-template/internal/repository"
	"github.com/phzeng0726/go-graphql-server-template/internal/service"

	_ "github.com/joho/godotenv/autoload"
)

func init() {
	config.InitConfig()
	if config.Env.Env == "" {
		log.Fatalln("Failed to load env")
	}
}

func main() {
	conn := database.Connect()
	database.SyncDatabase(conn)

	repos := repository.NewRepositories(conn)
	deps := service.Deps{
		Repos: repos,
	}
	services := service.NewServices(deps)

	handlers := delivery.NewGraphQLHandler(services)
	handlers.Init()

	addr := config.Env.Host + ":" + config.Env.Port
	log.Printf("Connect to http://%s/ for GraphQL playground", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
