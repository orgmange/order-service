package model

type OrderStatus string

const (
	OrderCreated  OrderStatus = "created"
	OrderPaid     OrderStatus = "paid"
	OrderCanceled OrderStatus = "canceled"
)

func (s OrderStatus) IsValid() bool {
	switch s {
	case OrderCanceled, OrderPaid, OrderCreated:
		return true
	default:
		return false
	}
}
