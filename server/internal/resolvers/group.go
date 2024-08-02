package resolvers

import (
	"context"

	"github.com/phzeng0726/go-graphql-server-template/internal/autogen"
	"github.com/phzeng0726/go-graphql-server-template/internal/domain"
)

type groupResolver struct {
	*Resolver
}

func (r *groupResolver) OwnerName(ctx context.Context, obj *domain.Group) (string, error) {
	user, err := r.Services.Users.GetUserById(ctx, obj.OwnerId)
	if err != nil {
		return "", err
	}

	return user.Name, nil
}

// Mutations
func (r *mutationResolver) CreateGroup(ctx context.Context, input autogen.NewGroupInput) (bool, error) {
	group := domain.Group{
		OwnerId: input.OwnerID,
		Remark:  input.Remark,
		EndedAt: input.EndedAt,
	}

	if err := r.Services.Groups.Create(ctx, group); err != nil {
		return false, err
	}

	return true, nil
}

/* EXAMPLE:
mutation createGroup(
  $ownerId: String! = "user0001"
  $remark: String! = "My first record",
  $endedAt: Timestamp! = "2024-01-10T12:34:56.123456789Z",
) {
  createGroup(
    input: {
      ownerId: $ownerId,
      remark: $remark,
      endedAt: $endedAt
    }
  )
}
*/

// Queries
func (r *queryResolver) GroupDetail(ctx context.Context, filter autogen.GroupDetailFilter) (*domain.Group, error) {
	group, err := r.Services.Groups.GetGroup(ctx, filter.GroupID)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

/* EXAMPLE:
query groupDetail(
  $groupId: Int! = 1
) {
  result: groupDetail(
    filter: {
      groupId: $groupId,
    }
  ) {
    id
    ownerName
    remark
    createdAt
    endedAt
  }
}
*/
