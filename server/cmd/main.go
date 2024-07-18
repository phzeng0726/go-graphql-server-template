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
}

func main() {
	conn := database.Connect()
	database.SyncDatabase(conn)

	repos := repository.NewRepositories(conn)
	srv := handler.NewDefaultServer(autogen.NewExecutableSchema(resolvers.NewConfig(repos)))

	// Only enable playground in development environment
	if config.Env.Env == "development" {
		http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	}

	// When not using the playground, handle CORS issues
	http.Handle("/graphql", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Content-Type", "application/json")
		srv.ServeHTTP(w, r)
	}))

	// Start the server with the specified Host and Port
	addr := config.Env.Host + ":" + config.Env.Port
	log.Printf("Connect to http://%s/ for GraphQL playground", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
