package resolvers

import (
	"context"

	"github.com/phzeng0726/go-graphql-server-template/internal/domain"
)

type groupResolver struct {
	*Resolver
}

// TODO 用 OwnerId 去查OwnerName
func (r *groupResolver) OwnerName(ctx context.Context, obj *domain.Group) (string, error) {
	return "ownerName", nil
}
