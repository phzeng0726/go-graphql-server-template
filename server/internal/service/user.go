package service

import (
	"context"

	"github.com/phzeng0726/go-graphql-server-template/internal/domain"
	"github.com/phzeng0726/go-graphql-server-template/internal/repository"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{
		repo: repo,
	}
}

func (s *UsersService) Create(ctx context.Context, user domain.User) error {
	if err := s.repo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *UsersService) GetUserById(ctx context.Context, userId string) (domain.User, error) {
	group, err := s.repo.GetUserById(ctx, userId)
	if err != nil {
		return domain.User{}, err
	}

	return group, nil
}
