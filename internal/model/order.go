package model

import (
	"fmt"
	"time"
)

type Order struct {
	id        uint
	creatorID uint
	status    OrderStatus
	createdAt time.Time
}

func (o *Order) GetCreatorID() uint {
	return o.creatorID
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
