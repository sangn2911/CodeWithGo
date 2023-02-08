package service

import (
	"context"
	"database/sql"
	"fmt"
	"go-research/internal/pkg/util"
	"go-research/internal/user/port"

	"gorm.io/gorm"
)

type UserService interface {
	GetUsersByName(context.Context, string, int, int) (util.ListResponse, error)
}

func NewUserService(db *gorm.DB, repo port.UserRepository) UserService {
	return &userService{
		db:   db,
		repo: repo,
	}
}

type userService struct {
	db   *gorm.DB
	repo port.UserRepository
}

func (u *userService) GetUsersByName(ctx context.Context, name string, pageSize, pageIdx int) (util.ListResponse, error) {
	userList, total, err := u.repo.GetUsersByName(ctx, name, pageSize, pageIdx)
	if err == sql.ErrNoRows {
		return util.ListResponse{}, fmt.Errorf("not found users by name %s", name)
	}
	if err != nil {
		return util.ListResponse{}, err
	}
	return util.ListResponse{
		Data: userList,
		Pagin: util.Pagination{
			PageSize: pageSize,
			PageIdx:  pageIdx,
			Total:    total,
		},
	}, nil
}
