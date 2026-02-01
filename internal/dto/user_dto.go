package dto

import "github.com/orgmange/order-service/internal/model"

type CreateUserRequest struct {
	Name  string
	Email string
}

type UpdateUserRequest struct {
	Name  string
	Email string
}

type UserResponse struct {
	ID    int
	Name  string
	Email string
}

func ToUser(req *CreateUserRequest) (*model.User, error) {
	return model.NewUser(0, req.Name, req.Email)
}

func ToResponse(user *model.User) *UserResponse {
	return &UserResponse{
		ID:    user.GetID(),
		Name:  user.GetName(),
		Email: user.GetEmail(),
	}
}
