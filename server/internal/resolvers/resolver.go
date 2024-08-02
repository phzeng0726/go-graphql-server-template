package resolvers

import (
	"github.com/phzeng0726/go-graphql-server-template/internal/autogen"
	"github.com/phzeng0726/go-graphql-server-template/internal/service"
)

type contextKey int

const (
	ContextKeyAuth contextKey = iota
	ContextKeyDB
)

type Resolver struct {
	Services *service.Services
}

func NewConfig(services *service.Services) autogen.Config {
	return autogen.Config{
		Resolvers: &Resolver{
			Services: services,
		},
	}
}

func (r *Resolver) Query() autogen.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Mutation() autogen.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Group() autogen.GroupResolver {
	return &groupResolver{r}
}
