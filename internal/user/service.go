package user

import (
	"context"
)

type Service struct {
	rep Repository
}

func NewService(rep Repository) *Service {
	return &Service{rep: rep}
}

// Create implements [UserService].
func (u *Service) Create(ctx context.Context, req *CreateParam) (*Model, error) {
	user, err := req.ToModel()
	if err != nil {
		return nil, err
	}
	savedUser, err := u.rep.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return savedUser, nil
}

// Delete implements [UserService].
func (u *Service) Delete(ctx context.Context, id uint) error {
	return u.rep.Delete(ctx, id)
}

// GetUser implements [UserService].
func (u *Service) GetUser(ctx context.Context, id uint) (*Model, error) {
	user, err := u.rep.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update implements [UserService].
func (u *Service) Update(ctx context.Context, id uint, req *UpdateParam) (*Model, error) {
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

	updatedUser, err := NewModel(id, name, email)
	if err != nil {
		return nil, err
	}

	return user, u.rep.Update(ctx, updatedUser)
}
