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

func (r *mutationResolver) CreateUser(ctx context.Context, input autogen.NewUserInput) (bool, error) {
	user := domain.User{
		Id:   input.UserID,
		Name: input.Username,
	}

	if err := r.Repos.Users.Create(ctx, user); err != nil {
		return false, err
	}

	return true, nil
}

/* EXAMPLE:
mutation createUser(
  $userId: String! = "user0001"
  $username: String! = "Rita Zeng",
) {
  createUser(
    input: {
      userId: $userId,
      username: $username,
    }
  )
}
*/
