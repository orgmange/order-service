package service

import (
	"github.com/orgmange/order-service/internal/dto"
	"github.com/orgmange/order-service/internal/model"
	"github.com/orgmange/order-service/internal/repository"
)

type UserServiceImpl struct {
	rep repository.UserRepository
}

func NewUserService(rep repository.UserRepository) UserService {
	return &UserServiceImpl{rep: rep}
}

// CreateUser implements [UserService].
func (u *UserServiceImpl) CreateUser(req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	user, err := dto.ToUser(req)
	if err != nil {
		return nil, err
	}
	savedUser, err := u.rep.Create(user)
	if err != nil {
		return nil, err
	}
	return dto.ToResponse(savedUser), nil
}

// DeleteUser implements [UserService].
func (u *UserServiceImpl) DeleteUser(id int) error {
	return u.rep.Delete(id)
}

// GetUser implements [UserService].
func (u *UserServiceImpl) GetUser(id int) (*dto.UserResponse, error) {
	user, err := u.rep.Get(id)
	if err != nil {
		return nil, err
	}
	return dto.ToResponse(user), nil
}

// UpdateUser implements [UserService].
func (u *UserServiceImpl) UpdateUser(id int, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := u.rep.Get(id)
	if err != nil {
		return nil, err
	}
	name := user.GetName()
	email := user.GetEmail()
	if req.Name != "" {
		name = req.Name
	}

	if req.Email != "" {
		email = req.Email
	}

	updatedUser, err := model.NewUser(id, name, email)
	if err != nil {
		return nil, err
	}

	return dto.ToResponse(user), u.rep.Update(updatedUser)
}
