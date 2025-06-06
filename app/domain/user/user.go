package user

import (
	"time"

	"github.com/tommylike/apaxrag/app/domain/pagination"
)

type User struct {
	ID        string
	UserName  string
	RealName  string
	Password  string
	Email     *string
	Phone     *string
	Status    int
	Creator   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Users []*User

type QueryParams struct {
	PaginationParam pagination.Param
	OrderFields     pagination.OrderFields
	UserName        string
	QueryValue      string
	Status          int
}
