package repository

import (
	"context"

	"github.com/phzeng0726/go-graphql-server-template/internal/domain"
	"gorm.io/gorm"
)

type Users interface {
	Create(ctx context.Context, user domain.User) error
	GetUserById(ctx context.Context, userId string) (domain.User, error)
}

type Groups interface {
	Create(ctx context.Context, group domain.Group) error
	GetGroup(ctx context.Context, groupId int) (domain.Group, error)
}

type Repositories struct {
	Groups Groups
	Users  Users
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Groups: NewGroupsRepo(db),
		Users:  NewUsersRepo(db),
	}
}
