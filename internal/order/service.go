package order

import (
	"context"
)

type Service struct {
	rep Repository
}

func NewService(rep Repository) *Service {
	return &Service{rep: rep}
}

func (u *Service) Create(ctx context.Context, req *CreateParam) (*Model, error) {
	order, err := req.ToModel()
	if err != nil {
		return nil, err
	}
	savedorder, err := u.rep.Create(ctx, order)
	if err != nil {
		return nil, err
	}
	return savedorder, nil
}

func (u *Service) Delete(ctx context.Context, id uint) error {
	return u.rep.Delete(ctx, id)
}

func (u *Service) Get(ctx context.Context, id uint) (*Model, error) {
	order, err := u.rep.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (u *Service) Update(ctx context.Context, id uint, req *UpdateParam) (*Model, error) {
	order, err := u.rep.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	creatorID := order.creatorID
	if req.CreatorID != 0 {
		creatorID = req.CreatorID
	}

	order.creatorID = creatorID
	return order, u.rep.Update(ctx, order)
}
