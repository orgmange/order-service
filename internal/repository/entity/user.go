package entity

import (
	"github.com/orgmange/order-service/internal/model"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func (u *User) ToModel() (*model.User, error) {
	return model.NewUser(u.ID, u.Name, u.Email)
}
