package service

import (
	"context"

	"github.com/phzeng0726/go-graphql-server-template/internal/domain"
	"github.com/phzeng0726/go-graphql-server-template/internal/repository"
)

type Users interface {
	Create(ctx context.Context, user domain.User) error
	GetUserById(ctx context.Context, userId string) (domain.User, error)
}

type Groups interface {
	Create(ctx context.Context, group domain.Group) error
	GetGroup(ctx context.Context, groupId int) (domain.Group, error)
}

type Services struct {
	Users  Users
	Groups Groups
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	usersService := NewUsersService(deps.Repos.Users)
	groupsService := NewGroupsService(deps.Repos.Groups)

	return &Services{
		Users:  usersService,
		Groups: groupsService,
	}
}
