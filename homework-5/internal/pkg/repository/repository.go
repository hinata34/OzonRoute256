package repository

import (
	"context"
	"errors"
)

var (
	ErrObjectNotFound = errors.New("object not found")
)

type UserRepo interface {
	Add(ctx context.Context, user *User) (int64, error)
	GetById(ctx context.Context, id int64) (*User, error)
	List(ctx context.Context) ([]*User, error)
	Update(ctx context.Context, user *User) (bool, error)
}

type AddressRepo interface {
	Add(ctx context.Context, address *Address) (int64, error)
	GetById(ctx context.Context, number int64) (*Address, error)
	List(ctx context.Context) ([]*Address, error)
	Update(ctx context.Context, address *Address) (bool, error)
}
