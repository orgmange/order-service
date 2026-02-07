package entity

import (
	"github.com/orgmange/order-service/internal/model"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CreatorID uint `gorm:"not null"`
	Status    string
}

func (o *Order) ToModel() (*model.Order, error) {
	return model.NewOrder(
		o.ID,
		o.CreatorID,
		model.OrderStatus(o.Status),
		o.CreatedAt)
}
