package user

import (
	"gorm.io/gorm"
)

type Entity struct {
	gorm.Model
	Name  string
	Email string
}

func (u *Entity) ToModel() (*Model, error) {
	return NewModel(u.ID, u.Name, u.Email)
}
