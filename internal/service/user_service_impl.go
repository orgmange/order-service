package service

import (
	"context"

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
func (u *UserServiceImpl) CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	user, err := dto.ToUser(req)
	if err != nil {
		return nil, err
	}
	savedUser, err := u.rep.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return dto.ToResponse(savedUser), nil
}

// DeleteUser implements [UserService].
func (u *UserServiceImpl) DeleteUser(ctx context.Context, id uint) error {
	return u.rep.Delete(ctx, id)
}

// GetUser implements [UserService].
func (u *UserServiceImpl) GetUser(ctx context.Context, id uint) (*dto.UserResponse, error) {
	user, err := u.rep.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ToResponse(user), nil
}

// UpdateUser implements [UserService].
func (u *UserServiceImpl) UpdateUser(ctx context.Context, id uint, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := u.rep.Get(ctx, id)
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

	return dto.ToResponse(user), u.rep.Update(ctx, updatedUser)
}
