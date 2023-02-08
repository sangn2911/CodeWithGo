package repository

import (
	"context"
	"go-research/internal/user/domain"
	"go-research/internal/user/port"

	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) port.UserRepository {
	return &UserRepository{
		db: db,
	}
}

type UserRepository struct {
	db *gorm.DB
}

func (u *UserRepository) GetUsersByName(ctx context.Context, name string, pageSize, pageIdx int) ([]domain.User, int64, error) {
	var count int64
	userlist := []domain.User{}
	if err := u.db.Model(domain.User{}).Where(domain.User{Name: name}).Limit(pageSize).Offset(pageIdx * pageSize).Find(&userlist).Error; err != nil {
		return userlist, 0, err
	}
	if err := u.db.Model(domain.User{}).Where(domain.User{Name: name}).Count(&count).Error; err != nil {
		return userlist, 0, err
	}
	return userlist, count, nil
}
