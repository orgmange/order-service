package dto

import (
	"time"

	"github.com/orgmange/order-service/internal/model"
)

type CreateOrderRequest struct {
	CreatorID uint
}

type UpdateOrderRequest struct {
	CreatorID uint
}

type OrderResponse struct {
	ID        uint
	CreatorID uint
	Status    string
	UpdatedAt time.Time
	CreatedAt time.Time
}

func (req *CreateOrderRequest) ToUser() (*model.Order, error) {
	return model.NewOrder(0, req.CreatorID, "", time.Time{})
}

func ToOrderResponse(order *model.Order) *OrderResponse {
	return &OrderResponse{
		ID:        order.GetID(),
		CreatorID: order.GetCreatorID(),
		Status:    string(order.GetStatus()),
		UpdatedAt: order.GetUpdatedAt(),
		CreatedAt: order.GetCreatedAt(),
	}
}
