package service

import "github.com/orgmange/order-service/internal/dto"

type UserService interface {
	CreateUser(req *dto.CreateUserRequest) (*dto.UserResponse, error)
	GetUser(id int) (*dto.UserResponse, error)
	UpdateUser(id int, req *dto.UpdateUserRequest) (*dto.UserResponse, error)
	DeleteUser(id int) error
}
