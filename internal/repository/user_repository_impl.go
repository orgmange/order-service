package repository

import (
	"fmt"
	"math/rand/v2"

	"github.com/orgmange/order-service/internal/model"
)

type UserRepositoryImpl struct {
	memoryDB map[int]model.User
}

// Create implements [UserRepository].
func (u *UserRepositoryImpl) Create(user *model.User) (*model.User, error) {
	ID := u.generateID()
	user.SetID(ID)
	u.memoryDB[ID] = *user
	return user, nil
}

// Delete implements [UserRepository].
func (u *UserRepositoryImpl) Delete(ID int) error {
	if _, ok := u.memoryDB[ID]; !ok {
		return fmt.Errorf("user not found")
	}

	delete(u.memoryDB, ID)
	return nil
}

// Get implements [UserRepository].
func (u *UserRepositoryImpl) Get(ID int) (*model.User, error) {
	user, ok := u.memoryDB[ID]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}

// Update implements [UserRepository].
func (u *UserRepositoryImpl) Update(user *model.User) error {
	u.memoryDB[user.GetID()] = *user
	return nil
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{
		memoryDB: make(map[int]model.User),
	}
}

func (u *UserRepositoryImpl) generateID() int {
	for {
		ID := rand.Int()
		_, ok := u.memoryDB[ID]
		if !ok {
			return ID
		}
	}
}
