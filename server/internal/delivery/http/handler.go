package http

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/phzeng0726/go-graphql-server-template/internal/autogen"
	"github.com/phzeng0726/go-graphql-server-template/internal/config"
	"github.com/phzeng0726/go-graphql-server-template/internal/resolvers"
	"github.com/phzeng0726/go-graphql-server-template/internal/service"
)

type GraphQLHandler struct {
	services *service.Services
}

func NewGraphQLHandler(
	services *service.Services,
) *GraphQLHandler {
	return &GraphQLHandler{
		services: services,
	}
}

func (h *GraphQLHandler) Init() {
	srv := handler.NewDefaultServer(autogen.NewExecutableSchema(resolvers.NewConfig(h.services)))

	if config.Env.Env == "development" {
		h.initPlayground()
	}

	h.initAPI(srv)
}

func (h *GraphQLHandler) initPlayground() {
	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
}

func (h *GraphQLHandler) initAPI(srv http.Handler) {
	http.Handle("/graphql", corsMiddleware(srv))
}
