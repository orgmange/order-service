package repository

import "github.com/orgmange/order-service/internal/model"

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	Get(ID int) (*model.User, error)
	Update(user *model.User) error
	Delete(ID int) error
}
