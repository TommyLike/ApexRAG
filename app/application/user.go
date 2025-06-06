package application

import (
	"context"
	"github.com/tommylike/apaxrag/app/domain/pagination"

	"github.com/tommylike/apaxrag/app/domain/user"

	"github.com/tommylike/apaxrag/app/domain/errors"
	"github.com/tommylike/apaxrag/pkg/util/hash"
	"github.com/tommylike/apaxrag/pkg/util/uuid"
)

type User interface {
	Query(ctx context.Context, params user.QueryParams) (user.Users, *pagination.Pagination, error)
	Get(ctx context.Context, id string) (*user.User, error)
	Create(ctx context.Context, item *user.User, roleIDs []string) (string, error)
	Update(ctx context.Context, id string, item *user.User, roleIDs []string) error
	Delete(ctx context.Context, id string) error
	UpdateStatus(ctx context.Context, id string, status int) error
}

func NewUser(
	userService user.Service,
) User {
	return &userApp{
		userService: userService,
	}
}

type userApp struct {
	userService user.Service
}

func (a *userApp) Query(ctx context.Context, params user.QueryParams) (user.Users, *pagination.Pagination, error) {
	result, pr, err := a.userService.Query(ctx, params)
	if err != nil {
		return nil, nil, err
	}
	return result, pr, nil
}

func (a *userApp) Get(ctx context.Context, id string) (*user.User, error) {
	item, err := a.userService.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *userApp) Create(ctx context.Context, item *user.User, roleIDs []string) (string, error) {
	err := a.checkUserName(ctx, item)
	if err != nil {
		return "", err
	}

	item.Password = hash.SHA1String(item.Password)
	item.ID = uuid.MustString()
	err = a.userService.Create(ctx, item)
	if err != nil {
		return "", err
	}

	return item.ID, nil
}

func (a *userApp) checkUserName(ctx context.Context, item *user.User) error {
	_, pr, err := a.userService.Query(ctx, user.QueryParams{
		PaginationParam: pagination.Param{OnlyCount: true},
		UserName:        item.UserName,
	})
	if err != nil {
		return err
	}
	if pr.Total > 0 {
		return errors.New400Response("The user name already exists")
	}
	return nil
}

func (a *userApp) Update(ctx context.Context, id string, item *user.User, roleIDs []string) error {
	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	}
	if oldItem == nil {
		return errors.ErrNotFound
	}
	if oldItem.UserName != item.UserName {
		err := a.checkUserName(ctx, item)
		if err != nil {
			return err
		}
	}

	if item.Password != "" {
		item.Password = hash.SHA1String(item.Password)
	} else {
		item.Password = oldItem.Password
	}

	item.ID = oldItem.ID
	item.Creator = oldItem.Creator
	item.CreatedAt = oldItem.CreatedAt
	return a.userService.Update(ctx, id, item)
}

func (a *userApp) Delete(ctx context.Context, id string) error {
	oldItem, err := a.userService.Get(ctx, id)
	if err != nil {
		return err
	}
	if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.userService.Delete(ctx, id)
}

func (a *userApp) UpdateStatus(ctx context.Context, id string, status int) error {
	oldItem, err := a.userService.Get(ctx, id)
	if err != nil {
		return err
	}
	if oldItem == nil {
		return errors.ErrNotFound
	}
	oldItem.Status = status

	err = a.userService.UpdateStatus(ctx, id, status)
	if err != nil {
		return err
	}

	return nil
}
