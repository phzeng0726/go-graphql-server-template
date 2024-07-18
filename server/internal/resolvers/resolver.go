package resolvers

import (
	"github.com/phzeng0726/go-graphql-server-template/internal/autogen"
	"github.com/phzeng0726/go-graphql-server-template/internal/repository"
)

type contextKey int

const (
	ContextKeyAuth contextKey = iota
	ContextKeyDB
)

type Resolver struct {
	Repos *repository.Repositories
}

func NewConfig(repos *repository.Repositories) autogen.Config {
	return autogen.Config{
		Resolvers: &Resolver{
			Repos: repos,
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
