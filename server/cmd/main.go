package main

import (
	"log"
	"net/http"

	"github.com/phzeng0726/go-graphql-server-template/internal/autogen"
	"github.com/phzeng0726/go-graphql-server-template/internal/config"
	"github.com/phzeng0726/go-graphql-server-template/internal/database"
	"github.com/phzeng0726/go-graphql-server-template/internal/repository"
	"github.com/phzeng0726/go-graphql-server-template/internal/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	_ "github.com/joho/godotenv/autoload"
)

func init() {
	config.InitConfig()
	if config.Env.Env == "" {
		log.Fatalln("Failed to load env")
	}

	log.Printf("Disable log: %v", config.Env.DisableLog)
}

func main() {
	conn := database.Connect()
	database.SyncDatabase(conn)

	repos := repository.NewRepositories(conn)

	srv := handler.NewDefaultServer(autogen.NewExecutableSchema(resolvers.NewConfig(repos)))

	// corsMiddleware
	graphqlHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Content-Type", "application/json")
		srv.ServeHTTP(w, r)
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", graphqlHandler) // 非透過playground的話，需要處理cors問題

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Env.Port)
	log.Fatal(http.ListenAndServe(":"+config.Env.Port, nil))
}
