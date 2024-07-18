package resolvers

import (
	"context"

	"github.com/phzeng0726/go-graphql-server-template/internal/autogen"
	"github.com/phzeng0726/go-graphql-server-template/internal/domain"
)

type mutationResolver struct {
	*Resolver
}

func (r *mutationResolver) CreateGroup(ctx context.Context, input autogen.NewGroupInput) (bool, error) {
	group := domain.Group{
		OwnerId: input.OwnerID,
		Remark:  input.Remark,
		EndedAt: input.EndedAt,
	}

	if err := r.Repos.Groups.Create(ctx, group); err != nil {
		return false, err
	}

	return true, nil
}

/*
mutation createGroup(
  $ownerId: String! = "aaa"
  $remark: String! = "remark",
  $products: [NewProductInput!] = [
        {
          name: "雞排",
          price: 100
        },
    		{
          name: "汽水",
          price: 30
        }
  ],
  $endedAt: Timestamp! = "2024-01-10T12:34:56.123456789Z",
) {
  createGroup(
    input: {
      ownerId: $ownerId,
      remark: $remark,
      products: $products
      endedAt: $endedAt
    }
  )
}
*/
