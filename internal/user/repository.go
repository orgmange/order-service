package user

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, user *Model) (*Model, error)
	Get(ctx context.Context, ID uint) (*Model, error)
	Update(ctx context.Context, user *Model) error
	Delete(ctx context.Context, ID uint) error
}
