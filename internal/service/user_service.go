package service

import (
	"context"

	"github.com/orgmange/order-service/internal/dto"
)

type UserService interface {
	CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error)
	GetUser(ctx context.Context, id uint) (*dto.UserResponse, error)
	UpdateUser(ctx context.Context, id uint, req *dto.UpdateUserRequest) (*dto.UserResponse, error)
	DeleteUser(ctx context.Context, id uint) error
}
