package model

import (
	"fmt"
	"time"
)

type Order struct {
	id        uint
	creatorID uint
	status    OrderStatus
	updatedAt time.Time
	createdAt time.Time
}

func (o *Order) GetID() uint {
	return o.id
}

func (o *Order) GetCreatorID() uint {
	return o.creatorID
}

func (o *Order) GetStatus() OrderStatus {
	return o.status
}

func (o *Order) GetUpdatedAt() time.Time {
	return o.updatedAt
}

func (o *Order) GetCreatedAt() time.Time {
	return o.createdAt
}

func NewOrder(id uint, creatorID uint, status OrderStatus, createdAt time.Time) (*Order, error) {
	if !status.IsValid() {
		return nil, fmt.Errorf("status not valid")
	}
	return &Order{
		id:        id,
		creatorID: creatorID,
		status:    status,
		createdAt: createdAt,
	}, nil
}
