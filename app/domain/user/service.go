package user

import (
	"context"
	"github.com/tommylike/apaxrag/app/domain/pagination"
)

func NewService(
	userRepo Repository,
) Service {
	return &service{
		userRepo: userRepo,
	}
}

type Service interface {
	Query(ctx context.Context, params QueryParams) (Users, *pagination.Pagination, error)
	Get(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, item *User) error
	Update(ctx context.Context, id string, item *User) error
	Delete(ctx context.Context, id string) error
	UpdateStatus(ctx context.Context, id string, status int) error
	UpdatePassword(ctx context.Context, id, password string) error
}

type service struct {
	userRepo Repository
}

func (s service) Query(ctx context.Context, params QueryParams) (Users, *pagination.Pagination, error) {
	return s.userRepo.Query(ctx, params)
}

func (s service) Get(ctx context.Context, id string) (*User, error) {
	return s.userRepo.Get(ctx, id)
}

func (s service) Create(ctx context.Context, item *User) error {
	return s.userRepo.Create(ctx, item)
}

func (s service) Update(ctx context.Context, id string, item *User) error {
	return s.userRepo.Update(ctx, id, item)
}

func (s service) Delete(ctx context.Context, id string) error {
	return s.userRepo.Delete(ctx, id)
}

func (s service) UpdateStatus(ctx context.Context, id string, status int) error {
	return s.userRepo.UpdateStatus(ctx, id, status)
}

func (s service) UpdatePassword(ctx context.Context, id, password string) error {
	return s.userRepo.UpdatePassword(ctx, id, password)
}
