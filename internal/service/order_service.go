package service

import (
	"context"

	"github.com/orgmange/order-service/internal/dto"
)

type OrderService interface {
	CreateOrder(ctx context.Context, req *dto.CreateOrderRequest) (*dto.OrderResponse, error)
	GetOrder(cxt context.Context, id uint) (*dto.OrderResponse, error)
	UpdateOrder(ctx context.Context, req *dto.UpdateOrderRequest) error
	DeleteOrder(ctx context.Context, id uint) error
}
