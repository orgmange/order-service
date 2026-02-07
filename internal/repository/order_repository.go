package repository

import (
	"context"

	"github.com/orgmange/order-service/internal/model"
)

type OrderRepository interface {
	Create(ctx context.Context, order *model.Order) (*model.Order, error)
	Get(ctx context.Context, id uint) (*model.Order, error)
	Update(ctx context.Context, order *model.Order) error
	Delete(ctx context.Context, id uint) error
}
