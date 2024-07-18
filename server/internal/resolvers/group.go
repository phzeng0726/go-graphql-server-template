package resolvers

import (
	"context"

	"github.com/phzeng0726/go-graphql-server-template/internal/domain"
)

type groupResolver struct {
	*Resolver
}

func (r *groupResolver) OwnerName(ctx context.Context, obj *domain.Group) (string, error) {
	user, err := r.Repos.Users.GetUserById(ctx, obj.OwnerId)
	if err != nil {
		return "", err
	}

	return user.Name, nil
}
