package resolvers

import (
	"context"

	"github.com/phzeng0726/go-graphql-server-template/internal/autogen"
	"github.com/phzeng0726/go-graphql-server-template/internal/domain"
)

// Mutations
func (r *mutationResolver) CreateUser(ctx context.Context, input autogen.NewUserInput) (bool, error) {
	user := domain.User{
		Id:   input.UserID,
		Name: input.Username,
	}

	if err := r.Services.Users.Create(ctx, user); err != nil {
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

// Queries
