package model

import (
	"fmt"

	"github.com/orgmange/order-service/internal/util"
)

type User struct {
	id    uint
	name  string
	email string
}

func NewUser(id uint, name string, email string) (*User, error) {
	if !util.NAME_REGEX.MatchString(name) {
		return nil, fmt.Errorf("not valid name")
	}

	if !util.EMAIL_REGEX.MatchString(email) {
		return nil, fmt.Errorf("not valid email")
	}

	return &User{
		id:    id,
		name:  name,
		email: email,
	}, nil
}

func (u *User) SetID(id uint) {
	u.id = id
}

func (u *User) GetID() uint {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}
