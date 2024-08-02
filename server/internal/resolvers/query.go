package resolvers

import (
	"context"
)

type queryResolver struct {
	*Resolver
}

func (r *queryResolver) Version(ctx context.Context) (string, error) {
	return "v0.0.1", nil
}
