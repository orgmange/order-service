package repository

import (
	"context"

	"github.com/orgmange/order-service/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Get(ctx context.Context, ID uint) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, ID uint) error
}
