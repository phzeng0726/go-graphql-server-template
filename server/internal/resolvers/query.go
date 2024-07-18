package resolvers

import (
	"context"

	"github.com/phzeng0726/go-graphql-server-template/internal/autogen"
	"github.com/phzeng0726/go-graphql-server-template/internal/domain"
)

type queryResolver struct {
	*Resolver
}

func (r *queryResolver) Version(ctx context.Context) (string, error) {
	return "v0.0.1", nil
}

func (r *queryResolver) GroupDetail(ctx context.Context, filter autogen.GroupDetailFilter) (*domain.Group, error) {
	group, err := r.Repos.Groups.GetGroup(ctx, filter.GroupID)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

/*
query groupDetail(
  $groupId: String = "H7YuFIi6rNPAe28CPiHP"
  $ownerId: String = "aaa"
) {
  result: groupDetail(
    filter: {
      groupId: $groupId,
      ownerId: $ownerId
    }
  ) {
    id
    ownerName
    remark
    createdAt
    closedAt
    products {
      id
      name
      price
    }
    orders {
      userId
      groupId
      items {
        productId
        quantity
      }
    }
  }
}
*/
