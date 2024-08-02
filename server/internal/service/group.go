package service

import (
	"context"

	"github.com/phzeng0726/go-graphql-server-template/internal/domain"
	"github.com/phzeng0726/go-graphql-server-template/internal/repository"
)

type GroupsService struct {
	repo repository.Groups
}

func NewGroupsService(repo repository.Groups) *GroupsService {
	return &GroupsService{
		repo: repo,
	}
}

func (s *GroupsService) Create(ctx context.Context, group domain.Group) error {
	if err := s.repo.Create(ctx, group); err != nil {
		return err
	}

	return nil
}

func (s *GroupsService) GetGroup(ctx context.Context, groupId int) (domain.Group, error) {
	group, err := s.repo.GetGroup(ctx, groupId)
	if err != nil {
		return domain.Group{}, err
	}

	return group, nil
}
