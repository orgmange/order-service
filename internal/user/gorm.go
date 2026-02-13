package user

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
func (u *gormRepository) Create(ctx context.Context, user *Model) (*Model, error) {
	userEntity := &Entity{
		Name:  user.GetName(),
		Email: user.GetEmail(),
	}
	err := gorm.G[Entity](u.db).Create(ctx, userEntity)
	if err != nil {
		return nil, err
	}

	return userEntity.ToModel()
}

// Delete implements [Repository].
func (u *gormRepository) Delete(ctx context.Context, ID uint) error {
	_, err := gorm.G[Entity](u.db).Where("id = ?", ID).Delete(ctx)
	return err
}

// Get implements [Repository].
func (u *gormRepository) Get(ctx context.Context, ID uint) (*Model, error) {
	userEntity, err := gorm.G[Entity](u.db).Where("id = ?", ID).First(ctx)
	if err != nil {
		return nil, err
	}
	return userEntity.ToModel()
}

// Update implements [Repository].
func (u *gormRepository) Update(ctx context.Context, user *Model) error {
	userEntity := Entity{}
	userEntity.Name = user.GetName()
	userEntity.Email = user.GetEmail()

	_, err := gorm.G[Entity](u.db).
		Where("id = ?", user.GetID()).
		Updates(ctx, userEntity)

	return err
}
