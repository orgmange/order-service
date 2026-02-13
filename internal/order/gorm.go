package order

import (
	"context"

	"gorm.io/gorm"
)

type gormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) Repository {
	return &gormRepository{
		db: db,
	}
}

// Create implements [Repository].
func (o *gormRepository) Create(ctx context.Context, order *Model) (*Model, error) {
	orderEntity := &Entity{
		CreatorID: order.GetCreatorID(),
	}
	err := gorm.G[Entity](o.db).
		Create(ctx, orderEntity)
	if err != nil {
		return nil, err
	}

	return orderEntity.ToModel()
}

// Delete implements [Repository].
func (o *gormRepository) Delete(ctx context.Context, id uint) error {
	_, err := gorm.G[Entity](o.db).
		Where("id = ?", id).
		Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Get implements [Repository].
func (o *gormRepository) Get(ctx context.Context, id uint) (*Model, error) {
	orderEntity, err := gorm.G[Entity](o.db).
		Where("id = ?", id).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return orderEntity.ToModel()
}

// Update implements [Repository].
func (o *gormRepository) Update(ctx context.Context, order *Model) error {
	orderEntity := Entity{
		CreatorID: order.GetCreatorID(),
	}
	_, err := gorm.G[Entity](o.db).
		Where("id = ", order.GetID()).
		Updates(ctx, orderEntity)

	return err
}
