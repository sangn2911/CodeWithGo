package port

import (
	"context"
	"go-research/internal/user/domain"
)

type UserRepository interface {
	GetUsersByName(context.Context, string, int, int) ([]domain.User, int64, error)
}
