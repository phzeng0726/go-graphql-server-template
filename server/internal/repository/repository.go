package repository

import (
	"context"

	"github.com/phzeng0726/go-graphql-server-template/internal/domain"
	"gorm.io/gorm"
)

type Groups interface {
	Create(ctx context.Context, group domain.Group) error
	GetGroup(ctx context.Context, groupId string) (domain.Group, error)
}

type Repositories struct {
	Groups Groups
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Groups: NewGroupsRepo(db),
	}
}
