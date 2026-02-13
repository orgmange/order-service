package order

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, order *Model) (*Model, error)
	Get(ctx context.Context, id uint) (*Model, error)
	Update(ctx context.Context, order *Model) error
	Delete(ctx context.Context, id uint) error
}
