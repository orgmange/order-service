package repository

import (
	"context"

	"github.com/orgmange/order-service/internal/model"
	"github.com/orgmange/order-service/internal/repository/entity"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{
		db: db,
	}
}

// Create implements [OrderRepository].
func (o *OrderRepositoryImpl) Create(ctx context.Context, order *model.Order) (*model.Order, error) {
	orderEntity := &entity.Order{
		CreatorID: order.GetCreatorID(),
	}
	err := gorm.G[entity.Order](o.db).Create(ctx, orderEntity)
	if err != nil {
		return nil, err
	}

	return orderEntity.ToModel()
}

// Delete implements [OrderRepository].
func (o *OrderRepositoryImpl) Delete(ctx context.Context, id uint) error {
	panic("unimplemented")
}

// Get implements [OrderRepository].
func (o *OrderRepositoryImpl) Get(ctx context.Context, id uint) (*model.Order, error) {
	panic("unimplemented")
}

// Update implements [OrderRepository].
func (o *OrderRepositoryImpl) Update(ctx context.Context, order *model.Order) error {
	panic("unimplemented")
}
