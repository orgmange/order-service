package repository

import (
	"context"

	"github.com/orgmange/order-service/internal/model"
	"github.com/orgmange/order-service/internal/repository/entity"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

// Create implements [UserRepository].
func (u *UserRepositoryImpl) Create(ctx context.Context, user *model.User) (*model.User, error) {
	userEntity := &entity.User{
		Name:  user.GetName(),
		Email: user.GetEmail(),
	}
	err := gorm.G[entity.User](u.db).Create(ctx, userEntity)
	if err != nil {
		return nil, err
	}

	return userEntity.ToModel()
}

// Delete implements [UserRepository].
func (u *UserRepositoryImpl) Delete(ctx context.Context, ID uint) error {
	_, err := gorm.G[entity.User](u.db).Where("id = ?", ID).Delete(ctx)
	return err
}

// Get implements [UserRepository].
func (u *UserRepositoryImpl) Get(ctx context.Context, ID uint) (*model.User, error) {
	userEntity, err := gorm.G[entity.User](u.db).Where("id = ?", ID).First(ctx)
	if err != nil {
		return nil, err
	}
	return userEntity.ToModel()
}

// Update implements [UserRepository].
func (u *UserRepositoryImpl) Update(ctx context.Context, user *model.User) error {
	var entity entity.User

	if err := u.db.First(&entity, user.GetID()).Error; err != nil {
		return err
	}

	entity.Name = user.GetName()
	entity.Email = user.GetEmail()

	return u.db.Save(&entity).Error
}
