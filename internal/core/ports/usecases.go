package ports

import (
	"context"

	domain "github.com/mitrovicsoftcoder/go-hexagonal-framework/internal/core/domain"
)

type UserUsecase interface {
	RegisterUser(ctx context.Context, user *domain.User) error
	FindByID(ctx context.Context, id string) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
}
