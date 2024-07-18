package repository

import (
	"context"

	"github.com/phzeng0726/go-graphql-server-template/internal/domain"
	"gorm.io/gorm"
)

type UsersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) Create(ctx context.Context, user domain.User) error {
	db := r.db.WithContext(ctx)

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) GetUserById(ctx context.Context, userId string) (domain.User, error) {
	var user domain.User
	db := r.db.WithContext(ctx)

	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
