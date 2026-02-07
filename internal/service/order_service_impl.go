package service

import (
	"context"
	"time"

	"github.com/orgmange/order-service/internal/dto"
	"github.com/orgmange/order-service/internal/model"
	"github.com/orgmange/order-service/internal/repository"
)

type OrderServiceImpl struct {
	rep repository.OrderRepository
}

func NewOrderService(rep repository.OrderRepository) OrderService {
	return &OrderServiceImpl{
		rep: rep,
	}
}

// CreateOrder implements [OrderService].
func (o *OrderServiceImpl) CreateOrder(ctx context.Context, req *dto.CreateOrderRequest) (*dto.OrderResponse, error) {
	order, err := req.ToUser()
	if err != nil {
		return nil, err
	}

	createdOrder, err := o.rep.Create(ctx, order)
	if err != nil {
		return nil, err
	}

	return dto.ToOrderResponse(createdOrder), nil
}

// DeleteOrder implements [OrderService].
func (o *OrderServiceImpl) DeleteOrder(ctx context.Context, id uint) error {
	return o.rep.Delete(ctx, id)
}

// GetOrder implements [OrderService].
func (o *OrderServiceImpl) GetOrder(ctx context.Context, id uint) (*dto.OrderResponse, error) {
	order, err := o.rep.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return dto.ToOrderResponse(order), nil
}

// UpdateOrder implements [OrderService].
func (o *OrderServiceImpl) UpdateOrder(ctx context.Context, req *dto.UpdateOrderRequest) error {
	order, err := model.NewOrder(0, req.CreatorID, "", time.Time{})
	if err != nil {
		return err
	}
	return o.rep.Update(ctx, order)
}
