package repository

import (
	"context"

	"github.com/phzeng0726/go-graphql-server-template/internal/domain"
	"gorm.io/gorm"
)

type GroupsRepo struct {
	db *gorm.DB
}

func NewGroupsRepo(db *gorm.DB) *GroupsRepo {
	return &GroupsRepo{
		db: db,
	}
}

func (r *GroupsRepo) Create(ctx context.Context, group domain.Group) error {
	db := r.db.WithContext(ctx)

	if err := db.Create(&group).Error; err != nil {
		return err
	}

	return nil
}

func (r *GroupsRepo) GetGroup(ctx context.Context, groupId string) (domain.Group, error) {
	return domain.Group{Id: "hi"}, nil
}
