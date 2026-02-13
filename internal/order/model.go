package order

import (
	"fmt"
	"time"
)

type Model struct {
	ID        uint
	CreatorID uint
	Status    Status
	UpdatedAt time.Time
	CreatedAt time.Time
}

func (m *Model) CheckValid() error {
	if !m.Status.IsValid() {
		return fmt.Errorf("status not valid")
	}
	return nil
}

type Status string

const (
	OrderCreated  Status = "created"
	OrderPaid     Status = "paid"
	OrderCanceled Status = "canceled"
)

func (s Status) IsValid() bool {
	switch s {
	case OrderCanceled, OrderPaid, OrderCreated:
		return true
	default:
		return false
	}
}
