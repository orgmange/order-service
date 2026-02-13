package order

import (
	"gorm.io/gorm"
)

type Entity struct {
	gorm.Model
	CreatorID uint `gorm:"not null"`
	Status    string
}

func (o *Entity) ToModel() *Model {
	return &Model{
		ID:        o.ID,
		CreatorID: o.CreatorID,
		Status:    Status(o.Status),
		UpdatedAt: o.UpdatedAt,
		CreatedAt: o.CreatedAt,
	}
}
